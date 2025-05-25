package main

import (
	"fmt"
	"time"
)

func main() {

	pingch := make(chan string)

	pongch := make(chan string)

	count := 5
	go ping(pingch, pongch, count)

	go pong(pingch, pongch, count)

	fmt.Println("starting ping pong game")
	pingch <- "ping 0"
	time.Sleep(15 * time.Second)
	fmt.Println("ping pong game complete")
}

func ping(pingch chan string, pongch chan string, count int) {

	for i := 0; i < count; i++ {

		msg := <-pingch
		fmt.Println("ping channel received", msg)
		time.Sleep(10 * time.Millisecond)
		pongch <- fmt.Sprintf("pong %d", i+1)

	}

}

func pong(pingch chan string, pongch chan string, count int) {

	for i := 0; i < count; i++ {

		msg := <-pongch
		fmt.Println("pong channel msg received", msg)
		time.Sleep(10 * time.Millisecond)
		pingch <- fmt.Sprintf("ping %d", i+1)

	}

}
