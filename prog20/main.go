package main

func main() {
	n := 45
	for i := 0; i < n; i++ {
		switch {
		case i%15 == 0:
			println("FizzBuzz")
		case i%3 == 0:
			println("Fizz")
		case i%5 == 0:
			println("Buzz")
		default:
			println(i)
		}
	}
}
