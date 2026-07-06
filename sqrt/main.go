package main

func main() {

	number := 4
	if number == 0 {
		println("Zero")
		return
	}
	count := 1
	for ; ; count++ {
		if count*count > number {
			println(count - 1)
			return
		}
	}

}
