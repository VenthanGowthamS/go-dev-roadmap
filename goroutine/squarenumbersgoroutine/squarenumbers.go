package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)

	go square(4, ch)

	fmt.Println("square number is", <-ch)

}

func square(number int, ch chan int) {

	ch <- number * number
}
