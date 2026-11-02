package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Menu struct {
	RestaurantID string
	Items        []string
}

// LoadMenus fetches menus concurrently with at most `workers` in flight,
// returning them in the same order as ids.
func LoadMenus(ids []string, workers int, getMenu func(id string) Menu) []Menu {
	results := make([]Menu, len(ids))
	workers = max(1, min(workers, len(ids)))

	jobs := make(chan int)
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				results[i] = getMenu(ids[i]) // each index written by exactly one worker
			}
		}()
	}

	for i := range ids {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return results
}

func main() {
	ids := make([]string, 10)
	for i := range ids {
		ids[i] = fmt.Sprintf("R%d", i+1)
	}
	fakeGetMenu := func(id string) Menu {
		time.Sleep(time.Duration(100+rand.Intn(400)) * time.Millisecond)
		return Menu{RestaurantID: id, Items: []string{"Biryani", "Dosa"}}
	}

	start := time.Now()
	menus := LoadMenus(ids, 3, fakeGetMenu)
	fmt.Printf("loaded %d menus in %v\n", len(menus), time.Since(start).Round(time.Millisecond))
	for _, m := range menus {
		fmt.Println(m.RestaurantID, m.Items)
	}
}
