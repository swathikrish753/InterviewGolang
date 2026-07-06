package main

func main() {
	number := 5
	fact := 1
	for i := 1; i <= number; i++ {
		fact *= i
	}
	println("the factorial of", number, "is", fact)
}
