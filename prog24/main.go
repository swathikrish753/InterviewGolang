package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}
type Result struct {
	JobID int
	Err   error
}

func main() {
	ctxVar := context.Background()
	ctx, cancel := context.WithDeadline(ctxVar, time.Now().Add(10*time.Second))
	defer cancel()
	WorkerPoolIntitialize(ctx, 5, 10)
}

func WorkerPoolIntitialize(ctx context.Context, workerCnt int, jobCount int) {
	jobChan := make(chan Job, jobCount)
	resultChan := make(chan Result, jobCount)
	wg := sync.WaitGroup{}
	for i := 0; i < workerCnt; i++ {
		wg.Add(1)
		go worker(ctx, jobChan, resultChan, &wg)
	}
	for j := 1; j <= jobCount; j++ {
		jobChan <- Job{ID: j}
	}
	close(jobChan)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for res := range resultChan {
		if res.Err != nil {
			fmt.Printf("Job %d failed with error: %v\n", res.JobID, res.Err)
		} else {
			fmt.Printf("Job %d completed successfully\n", res.JobID)
		}
	}

}

func worker(ctx context.Context, jobChan chan Job, resultChan chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobChan:
			if !ok {
				return
			}
			if job.ID%2 == 0 {
				resultChan <- Result{JobID: job.ID, Err: fmt.Errorf("job %d failed", job.ID)}
			} else {
				resultChan <- Result{JobID: job.ID, Err: nil}
			}
		case <-ctx.Done():
			fmt.Printf("Worker exiting due to context cancellation: %v\n", ctx.Err())
			return
		}
	}
}
