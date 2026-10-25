package main

func main() {
	channel := make(chan int)
	doneChannel := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			channel <- i
		}
		doneChannel <- true
	}()

	go func() {
		for i := 5; i < 10; i++ {
			channel <- i
		}
		doneChannel <- true
	}()

	go func() {
		<-doneChannel
		<-doneChannel
		close(channel)
	}()

	for num := range channel {
		println(num)
	}
}
