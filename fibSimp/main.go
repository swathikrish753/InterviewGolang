package main

func main() {
	number := 4

	a := 0
	b := 1

	for i := 0; i <= number; i++ {
		if i == 0 {
			println(a)
		} else if i == 1 {
			println(b)
		} else {
			a, b = b, a+b
			println(b)
		}
	}
}
