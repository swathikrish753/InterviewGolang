package main

func main() {
	n := 10
	println("the", n, "th Fibonacci number is", fib(n))
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
