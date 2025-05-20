package main

import (
	"fmt"
	"time"
)

func main() {

	messages := []string{"Mr.coder", "is", "consistent", "hard", "to", "break"}

	ch := make(chan string, 3) //we are making it buffered
	for _, msg := range messages {
		go pushtochannel(msg, ch)
	}

	for i := 0; i < len(messages); i++ {
		time.Sleep(time.Millisecond * 200)
		receiver := <-ch
		fmt.Println(receiver)
	}
}

func pushtochannel(msg string, ch chan string) {
	fmt.Println("sending message......", msg)
	ch <- msg
	fmt.Println("sent message--->", msg)

}
