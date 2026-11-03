package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Menu struct {
	RestaurantID string
	Items        []string
}

type Job struct {
	Index int
	ID    string
}

type Result struct {
	Index int
	ID    string
	Menu  Menu
	Err   error
}

func worker(getMenu func(string) (Menu, error), jobs <-chan Job, results chan<- Result, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		select {
		case <-done:
			return // a menu already failed; don't start new work
		default:
		}
		menu, err := getMenu(job.ID)
		results <- Result{Index: job.Index, ID: job.ID, Menu: menu, Err: err}
	}
}

func LoadMenus(ids []string, workers int, getMenu func(string) (Menu, error)) ([]Menu, error) {
	jobs := make(chan Job, len(ids))
	results := make(chan Result, len(ids)) // buffered so workers never block
	done := make(chan struct{})

	// 1. Start workers
	var wg sync.WaitGroup
	for range max(1, min(workers, len(ids))) {
		wg.Add(1)
		go worker(getMenu, jobs, results, done, &wg)
	}

	// 2. Send jobs
	for i, id := range ids {
		jobs <- Job{Index: i, ID: id}
	}
	close(jobs)

	// 3. Close results once all workers exit
	go func() {
		wg.Wait()
		close(results)
	}()

	// 4. Collect results
	menus := make([]Menu, len(ids))
	var firstErr error
	for r := range results {
		if r.Err != nil {
			if firstErr == nil {
				firstErr = fmt.Errorf("loading menu %s: %w", r.ID, r.Err)
				close(done)
			}
			continue
		}
		menus[r.Index] = r.Menu
	}

	if firstErr != nil {
		return nil, firstErr
	}
	return menus, nil
}

func main() {
	ids := make([]string, 10)
	for i := range ids {
		ids[i] = fmt.Sprintf("R%d", i+1)
	}

	var started atomic.Int32
	fakeGetMenu := func(id string) (Menu, error) {
		started.Add(1)
		time.Sleep(time.Duration(100+rand.Intn(400)) * time.Millisecond)
		if id == "R4" {
			return Menu{}, errors.New("server down")
		}
		return Menu{RestaurantID: id, Items: []string{"Biryani", "Dosa"}}, nil
	}

	start := time.Now()
	menus, err := LoadMenus(ids, 3, fakeGetMenu)
	fmt.Println("started:", started.Load(), "loaded:", len(menus), "err:", err, "took:", time.Since(start).Round(time.Millisecond))
}
