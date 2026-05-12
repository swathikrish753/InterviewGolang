package main

func main() {
	channel := make(chan int, 10)
	doneChannel := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(num int) {
			channel <- num
			doneChannel <- true
		}(i)

	}

	for i := 0; i < 10; i++ {
		<-doneChannel
	}
	close(channel)

	for num := range channel {
		println(num)
	}
}
