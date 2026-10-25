package main

import "fmt"

// Step 1: Define a Node (box that holds one digit + pointer to next box)

type Node struct {
	Val  int
	Next *Node
}

// Step 2: Add two linked lists digit by digit
func addTwoNumbers(l1 *Node, l2 *Node) *Node {

	carry := 0       // extra 1 to carry over when sum >= 10
	dummy := &Node{} // empty starting node (we won't use its Val)
	current := dummy // pointer to build our result list

	// Keep going as long as there's a digit in l1, l2, or a leftover carry
	for l1 != nil || l2 != nil || carry != 0 {

		sum := carry // start with carry from previous step

		// Add digit from l1 if available
		if l1 != nil {
			sum = sum + l1.Val
			l1 = l1.Next
		}

		// Add digit from l2 if available
		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}

		// Example: sum = 13  →  digit = 3, carry = 1
		digit := sum % 10 // remainder  (13 % 10 = 3)
		carry = sum / 10  // carry over (13 / 10 = 1)

		// Create a new node with the digit and attach it
		current.Next = &Node{Val: digit}
		current = current.Next
	}

	return dummy.Next // skip the empty starting node
}

// ---- Helper: slice → linked list ----
func makeList(digits []int) *Node {
	dummy := &Node{}
	cur := dummy
	for _, d := range digits {
		cur.Next = &Node{Val: d}
		cur = cur.Next
	}
	return dummy.Next
}

// ---- Helper: print linked list ----
func printList(node *Node) {
	for node != nil {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print(" -> ")
		}
		node = node.Next
	}
	fmt.Println()
}

func main() {
	// 342 + 465 = 807
	// Stored in reverse: [2->4->3] + [5->6->4] = [7->0->8]
	l1 := makeList([]int{2, 4, 3})
	l2 := makeList([]int{5, 6, 4})

	fmt.Print("List 1 : ")
	printList(l1)
	fmt.Print("List 2 : ")
	printList(l2)
	fmt.Print("Result : ")
	printList(addTwoNumbers(l1, l2))
}
