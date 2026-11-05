package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
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

var failID = "R4" // set to "" to switch off the failure

// fetchMenu pretends to call a slow server.
func fetchMenu(ctx context.Context, id string) (Menu, error) {
	delay := time.Duration(100+rand.Intn(400)) * time.Millisecond
	select {
	case <-time.After(delay):
	case <-ctx.Done():
		return Menu{}, ctx.Err()
	}
	if id == failID {
		return Menu{}, errors.New("server down")
	}
	return Menu{RestaurantID: id, Items: []string{"Biryani", "Dosa"}}, nil
}

func worker(ctx context.Context, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		if ctx.Err() != nil {
			return
		}
		menu, err := fetchMenu(ctx, job.ID)
		results <- Result{Index: job.Index, ID: job.ID, Menu: menu, Err: err}
	}
}

func LoadMenus(ctx context.Context, ids []string, workers int) ([]Menu, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	jobs := make(chan Job, len(ids))
	results := make(chan Result, len(ids))

	if workers <= 0 {
		workers = 1
	}

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(ctx, jobs, results, &wg)
	}

	for i, id := range ids {
		jobs <- Job{Index: i, ID: id}
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	menus := make([]Menu, len(ids))
	var firstErr error
	for r := range results {
		if r.Err != nil {
			if firstErr == nil {
				firstErr = fmt.Errorf("loading menu %s: %w", r.ID, r.Err)
				cancel() // stop in-flight fetches
			}
			continue
		}
		menus[r.Index] = r.Menu
	}

	if firstErr != nil {
		return nil, firstErr
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return menus, nil
}

func main() {
	ids := make([]string, 10)
	for i := range ids {
		ids[i] = fmt.Sprintf("R%d", i+1)
	}

	// Case 1: R4 fails
	start := time.Now()
	_, err := LoadMenus(context.Background(), ids, 3)
	fmt.Println("case 1:", err, "took:", time.Since(start).Round(time.Millisecond))

	// Case 2: timeout, no failures
	failID = ""
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	start = time.Now()
	_, err = LoadMenus(ctx, ids, 3)
	fmt.Println("case 2:", err, "| is deadline:", errors.Is(err, context.DeadlineExceeded), "took:", time.Since(start).Round(time.Millisecond))
}
