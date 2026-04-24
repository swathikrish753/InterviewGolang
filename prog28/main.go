package main

import "fmt"

func main() {
	// Print table header with column names
	fmt.Printf("%-20s %10s %10s\n", "Name", "Age", "Score")

	// Print data rows with formatted columns
	fmt.Printf("%-20s %10d %10.2f\n", "Alice", 30, 85.6)
	fmt.Printf("%-20s %10d %10.2f\n", "Bob", 24, 92.3)
	fmt.Printf("%-20s %10d %10.2f\n", "Charlie", 29, 88.1)
}
