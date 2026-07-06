package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// This struct owns the goroutine
type EmailService struct {
	wg     sync.WaitGroup
	cancel context.CancelFunc
}

// start launches the gorutine
func (e *EmailService) start(parent context.Context) {
	ctx, cancel := context.WithCancel(parent)
	e.cancel = cancel
	e.wg.Add(1)
	go func() {
		defer e.wg.Done()
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Email service shutting down")
				return

			case t := <-ticker.C:
				fmt.Println("Sending batch emails at", t)
			}
		}
	}()
}

// stop terminates the goroutine
func (e *EmailService) stop() {
	e.cancel()
}

// wait ensures cleanup
func (e *EmailService) wait() {
	e.wg.Wait()
}

func main() {
	// Option 1: manual cancellation
	//ctx := context.Background()

	// Option 2: auto-timeout (uncomment to test ctx.Done automatically)
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	EmailService := &EmailService{}
	EmailService.start(ctx)

	// Simulate doing other work for 5 seconds
	time.Sleep(5 * time.Second)

	fmt.Println("Deployment shutdown triggered")
	EmailService.stop()
	EmailService.wait()
}
