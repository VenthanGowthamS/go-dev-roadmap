package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)

	ch2 := make(chan string)

	//go routine and anaonymous function
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "received message from channel 1"
	}()

	//go routine and anaonymous function
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "received message from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received", msg2)

	}

}
