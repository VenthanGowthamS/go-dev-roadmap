package main

import (
	"fmt"
	"time"
)

func concurrency(name string) {
	for i := 0; i < 5; i++ {

		fmt.Println(name, ":", i)
		time.Sleep(1000 * time.Millisecond)
		// give time for goroutine to run
	}
}

func main() {
	go concurrency("Mr.Coder")

	go concurrency("Mr.Learner")

	time.Sleep(2000 * time.Millisecond)
	// give time for goroutine to run
	fmt.Println("main closed")
}
