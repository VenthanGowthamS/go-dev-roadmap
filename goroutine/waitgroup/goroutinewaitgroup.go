package main

import (
	"fmt"
	"sync"
)

func printmessage(msg string, wg *sync.WaitGroup) {
	fmt.Println(msg)
	wg.Done()

}

func main() {

	var wg sync.WaitGroup

	message := []string{"hello", "im", "learning", "program logic", "and", "go"}
	for _, msg := range message {

		wg.Add(1)
		go printmessage(msg, &wg)

	}
	wg.Wait()
}
