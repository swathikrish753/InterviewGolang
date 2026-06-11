package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

// ---- bucket ----------------------------------------------------------------

const (
	maxRequests    = 10
	windowDuration = 1 * time.Minute
)

// bucket holds the token count for a single IP.
// count is accessed atomically so reads/decrements don't need the global mutex.
type bucket struct {
	count int64 // remaining tokens this window
}

func newBucket() *bucket {
	return &bucket{count: maxRequests}
}

// allow tries to consume one token. Returns true if the request is permitted.
func (b *bucket) allow() bool {
	for {
		cur := atomic.LoadInt64(&b.count)
		if cur <= 0 {
			return false
		}
		// CAS: only decrement if the value we read is still current
		if atomic.CompareAndSwapInt64(&b.count, cur, cur-1) {
			return true
		}
		// another goroutine changed count between Load and CAS → retry
	}
}

// reset restores the bucket to a full window.
func (b *bucket) reset() {
	atomic.StoreInt64(&b.count, maxRequests)
}

// ---- limiter ---------------------------------------------------------------

// RateLimiter manages one bucket per IP and a background refill goroutine.
type RateLimiter struct {
	mu      sync.RWMutex
	buckets map[string]*bucket
	done    chan struct{}
}

func NewRateLimiter() *RateLimiter {
	rl := &RateLimiter{
		buckets: make(map[string]*bucket),
		done:    make(chan struct{}),
	}
	go rl.refillLoop()
	return rl
}

// refillLoop resets every bucket at the start of each new window.
func (rl *RateLimiter) refillLoop() {
	ticker := time.NewTicker(windowDuration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			rl.resetAll()
		case <-rl.done:
			return
		}
	}
}

func (rl *RateLimiter) resetAll() {
	rl.mu.RLock()
	defer rl.mu.RUnlock()

	for _, b := range rl.buckets {
		b.reset()
	}
	log.Printf("[limiter] all buckets reset (%d IPs tracked)", len(rl.buckets))
}

// getBucket returns the existing bucket for ip, or creates a new one.
func (rl *RateLimiter) getBucket(ip string) *bucket {
	// fast path: bucket already exists
	rl.mu.RLock()
	b, ok := rl.buckets[ip]
	rl.mu.RUnlock()
	if ok {
		return b
	}

	// slow path: first request from this IP
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// double-check after acquiring write lock (another goroutine may have
	// created it between our RUnlock and Lock)
	if b, ok = rl.buckets[ip]; ok {
		return b
	}

	b = newBucket()
	rl.buckets[ip] = b
	return b
}

// Allow is the public entry point; returns true if the IP may proceed.
func (rl *RateLimiter) Allow(ip string) bool {
	return rl.getBucket(ip).allow()
}

// Stop shuts down the background goroutine (call on server shutdown).
func (rl *RateLimiter) Stop() {
	close(rl.done)
}

// ---- middleware ------------------------------------------------------------

func rateLimitMiddleware(rl *RateLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr // in production: parse X-Forwarded-For

		if !rl.Allow(ip) {
			log.Printf("[blocked] ip=%s path=%s", ip, r.URL.Path)
			w.Header().Set("Retry-After", "60")
			http.Error(w, "429 Too Many Requests — rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ---- handlers & main -------------------------------------------------------

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %s\n", r.URL.Path)
}

func main() {
	rl := NewRateLimiter()
	defer rl.Stop()

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: rateLimitMiddleware(rl, mux),
	}

	log.Println("Server running on :8080  (limit: 10 req/IP/min)")
	log.Fatal(server.ListenAndServe())
}
