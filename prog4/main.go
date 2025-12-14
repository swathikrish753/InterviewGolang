package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {

	stack := []int{}
	for _, token := range tokens {
		switch token {
		case " + ", " - ", " * ", " / ":
			// Pop last two operands
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var res int
			switch token {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b // truncates toward zero
			}

			stack = append(stack, res)
		default:
			// Convert string to integer
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func main() {
	expr := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(expr)) // Output: 9
}
