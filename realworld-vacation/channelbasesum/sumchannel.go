package main

import "fmt"

func main() {

	ch := make(chan int)
	//defining integer un buffered channel to send to go routine

	//we are passing the channel created to go routine will be send to channel and main go routine will
	//receive from channel
	go calculatesum(ch)
	sum := 0

	//receive the values from channel and sum it then
	for values := range ch {
		sum += values
	}
	fmt.Println("sum is", sum)
}

func calculatesum(ch chan int) {
	for i := 0; i <= 5; i++ {
		ch <- i

		//send values to channel
	}
	close(ch) //close the channel after values being sent
}
