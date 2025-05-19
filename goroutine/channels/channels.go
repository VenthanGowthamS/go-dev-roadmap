package main

import (
	"fmt"
)

func main() {

	messages := []string{"hello", "baby", "how are you"}
	ch := make(chan string)
	for _, msg := range messages {
		go sendmessage(msg, ch)
	}
	fmt.Println(len(ch))
	for i := 0; i < len(messages); i++ {
		receivedmessage := <-ch
		fmt.Println("message are", receivedmessage)
	}
}

//chan string means "a channel that can carry string values."

// chan is the channel type keyword in Go.
func sendmessage(msg string, ch chan string) {
	ch <- msg
	//send the string message to channel
	fmt.Println("message sent to channel")

}
