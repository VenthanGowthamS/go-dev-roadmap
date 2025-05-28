package main

import (
	"fmt"
	"time"
)

func main() {

	go greet("Coder")

	go greet("Learner")

	time.Sleep(2 * time.Second)
}

func greet(name string) {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello is", name)
}
