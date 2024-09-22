package main

import "fmt"

func pingPong(notifier chan string, i int) {
	isPing := i%2 == 0

	if isPing {
		notifier <- "Ping"
	} else {
		notifier <- "Pong"
	}
}

func main() {
	numbersChannel := make(chan int)
	pingPongChannel := make(chan string)
	done := make(chan bool)

	go func() {
		for {
			job, ok := <-numbersChannel
			if ok {

				fmt.Println(job, ok)
			} else {
				fmt.Printf("Not ok %v", ok)
				done <- true
				return
			}

		}
	}()

	go func() {
		for {
			job, ok := <-pingPongChannel
			if ok {

				fmt.Println(job, ok)
			} else {
				fmt.Printf("Not ok %v", ok)
				done <- true
				return
			}

		}
	}()

	for i := 0; i <= 5; i++ {
		someOtherChannel := <-numbersChannel
		fmt.Println(someOtherChannel)
		numbersChannel <- i
		pingPong(pingPongChannel, i)
	}
	close(numbersChannel)

	<-done
}
