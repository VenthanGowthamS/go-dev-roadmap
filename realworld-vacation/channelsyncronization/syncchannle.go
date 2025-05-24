package main

import (
	"fmt"
	"time"
)

func main() {

	//declaring a bool channel in channelconst
	done := make(chan bool)
	fmt.Println("")
	//call the async function go routine and pass the channel
	go sendsignal_to_worker(done)
	fmt.Println("waitig for signal..")
	//wait till some bool value
	<-done

	fmt.Println("value received")

}

func sendsignal_to_worker(done chan bool) {
	fmt.Println("we recevied signal let us send to main go routine to proceed ")
	time.Sleep(2 * time.Second)

	done <- true
}
