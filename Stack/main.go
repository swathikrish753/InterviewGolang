package main

import (
	"fmt"
)

// Stack represents a simple LIFO (Last-In, First-Out) stack.
// The data is stored in a slice of integers.
type Stack struct {
	items []int
}

// Push adds an item to the top of the stack.

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack.
// It also checks if the stack is empty.
func (s *Stack) Pop() (int, bool) {
	// Check if the stack is empty
	if s.IsEmpty() {
		// Return 0 (default value for int) and false (indicating failure)
		return 0, false
	}

	// Get the index of the last element (the "top" of the stack)
	lastIndex := len(s.items) - 1

	// Get the item at the top
	item := s.items[lastIndex]

	// Remove the top item by slicing the slice up to the last element
	s.items = s.items[:lastIndex]

	// Return the item and true (indicating success)
	return item, true
}

// Peek returns the item at the top of the stack without removing it.
// It also checks if the stack is empty.
func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		// Return 0 and false
		return 0, false
	}

	// Get the index of the last element (the "top")
	lastIndex := len(s.items) - 1

	// Return the item and true
	return s.items[lastIndex], true
}

// IsEmpty checks if the stack contains any items.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	// 1. Create a new stack
	myStack := Stack{}

	fmt.Println("Is stack empty (initial)?", myStack.IsEmpty()) // true

	// 2. Push operations
	myStack.Push(10)
	myStack.Push(20)
	myStack.Push(30)
	fmt.Println("\nStack after PUSH(10, 20, 30):", myStack.items) // [10 20 30]

	// 3. Peek operation
	if val, ok := myStack.Peek(); ok {
		fmt.Printf("PEEK result: %d (Stack remains: %v)\n", val, myStack.items) // 30
	}

	// 4. Pop operations
	if val, ok := myStack.Pop(); ok {
		fmt.Printf("POP result: %d (Stack is now: %v)\n", val, myStack.items) // 30, [10 20]
	}
	if val, ok := myStack.Pop(); ok {
		fmt.Printf("POP result: %d (Stack is now: %v)\n", val, myStack.items) // 20, [10]
	}

	// 5. Final check
	fmt.Println("\nIs stack empty (after Pop)?", myStack.IsEmpty()) // false (10 is left)

	// 6. Pop the last item
	myStack.Pop()
	fmt.Println("Stack after final Pop:", myStack.items)      // []
	fmt.Println("Is stack empty (final)?", myStack.IsEmpty()) // true

	// 7. Pop on an empty stack (should fail gracefully)
	if _, ok := myStack.Pop(); !ok {
		fmt.Println("\nAttempted POP on an empty stack: Operation failed as expected.")
	}
}
