package main

import (
	"fmt"
	"time"
)

// Two goroutines pass a ball back and forth over unbuffered channels.
// Each send blocks until the other side receives, so they take strict turns.
func main() {
	pingch := make(chan string)
	pongch := make(chan string)
	done := make(chan struct{})

	rounds := 5
	go ping(pingch, pongch, rounds)
	go pong(pingch, pongch, rounds, done)

	fmt.Println("starting ping pong game")
	pingch <- "ping 0"

	// Wait until the last pong is played instead of sleeping for a guessed duration.
	<-done
	fmt.Println("ping pong game complete")
}

func ping(pingch <-chan string, pongch chan<- string, rounds int) {
	for i := 0; i < rounds; i++ {
		msg := <-pingch
		fmt.Println("ping channel received", msg)
		time.Sleep(10 * time.Millisecond)
		pongch <- fmt.Sprintf("pong %d", i+1)
	}
}

func pong(pingch chan<- string, pongch <-chan string, rounds int, done chan<- struct{}) {
	for i := 0; i < rounds; i++ {
		msg := <-pongch
		fmt.Println("pong channel msg received", msg)
		time.Sleep(10 * time.Millisecond)

		// The last round has no ping left to answer it. Sending would block forever
		// (a leaked goroutine), so signal that the game is over instead.
		if i == rounds-1 {
			close(done)
			return
		}
		pingch <- fmt.Sprintf("ping %d", i+1)
	}
}
