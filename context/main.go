package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Demonstrating Context Functions in Go ===")

	// 1. context.Background() - Root context, never canceled
	fmt.Println("\n1. context.Background() - Root context")
	ctxBg := context.Background()
	fmt.Printf("Background context: %v\n\n", ctxBg)

	// 2. context.TODO() - Placeholder for contexts not yet determined
	fmt.Println("2. context.TODO() - Placeholder context")
	ctxTodo := context.TODO()
	fmt.Printf("TODO context: %v\n\n", ctxTodo)

	// 3. context.WithCancel() - Creates a context that can be canceled
	fmt.Println("3. context.WithCancel() - Cancelable context")
	ctxCancel, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctxCancel.Done():
			fmt.Println("Goroutine canceled via WithCancel")
		case <-time.After(2 * time.Second):
			fmt.Println("Goroutine completed normally")
		}
	}()
	time.Sleep(1 * time.Second)
	cancel() // Cancel the context
	wg.Wait()
	fmt.Printf("Cancel context error: %v\n\n", ctxCancel.Err())

	// 4. context.WithDeadline() - Context with absolute deadline
	fmt.Println("4. context.WithDeadline() - Context with deadline")
	deadline := time.Now().Add(2 * time.Second)
	ctxDeadline, cancelDeadline := context.WithDeadline(context.Background(), deadline)
	defer cancelDeadline()
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctxDeadline.Done():
			fmt.Println("Goroutine canceled via WithDeadline")
		}
	}()
	time.Sleep(3 * time.Second) // Wait past deadline
	wg.Wait()
	fmt.Printf("Deadline context error: %v\n\n", ctxDeadline.Err())

	// 5. context.WithTimeout() - Context with relative timeout
	fmt.Println("5. context.WithTimeout() - Context with timeout")
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelTimeout()
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctxTimeout.Done():
			fmt.Println("Goroutine canceled via WithTimeout")
		}
	}()
	time.Sleep(3 * time.Second) // Wait past timeout
	wg.Wait()
	fmt.Printf("Timeout context error: %v\n\n", ctxTimeout.Err())

	// 6. context.WithValue() - Context with key-value pairs
	fmt.Println("6. context.WithValue() - Context with values")
	type keyType string
	userKey := keyType("user")
	ctxValue := context.WithValue(context.Background(), userKey, "john_doe")
	value := ctxValue.Value(userKey)
	fmt.Printf("Retrieved value from context: %v\n", value)

	// Nested contexts with values
	ctxNested := context.WithValue(ctxValue, "session", "abc123")
	sessionValue := ctxNested.Value("session")
	fmt.Printf("Retrieved nested value: %v\n\n", sessionValue)

	// 7. Context methods demonstration
	fmt.Println("7. Context Methods Demonstration")
	ctxMethods, cancelMethods := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelMethods()

	// Deadline() method
	if deadline, ok := ctxMethods.Deadline(); ok {
		fmt.Printf("Context deadline: %v\n", deadline)
	}

	// Done() channel
	fmt.Println("Waiting for context to be done...")
	select {
	case <-ctxMethods.Done():
		fmt.Println("Context is done")
	}

	// Err() method
	fmt.Printf("Context error: %v\n\n", ctxMethods.Err())

	// 8. Combining contexts - WithCancel + WithValue
	fmt.Println("8. Combining Contexts - WithCancel + WithValue")
	ctxCombined, cancelCombined := context.WithCancel(
		context.WithValue(context.Background(), "request_id", "req-123"))
	defer cancelCombined()

	wg.Add(1)
	go func() {
		defer wg.Done()
		reqID := ctxCombined.Value("request_id")
		fmt.Printf("Goroutine processing request: %v\n", reqID)
		select {
		case <-ctxCombined.Done():
			fmt.Println("Combined context canceled")
		case <-time.After(1 * time.Second):
			fmt.Println("Request processed")
		}
	}()
	time.Sleep(500 * time.Millisecond)
	cancelCombined()
	wg.Wait()

	fmt.Println("\n=== All Context Functions Demonstrated ===")
}
