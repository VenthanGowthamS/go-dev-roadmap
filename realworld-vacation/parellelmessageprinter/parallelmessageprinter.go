package main

import (
	"fmt"
)

//start multiple go routine to print

func main() {

	messages := []string{"go", "lang", "learning", "is", "fun", "making", "it", "consistent", "is", "happy"}

	for _, v := range messages {
		go printmsg(v)
	}
	fmt.Scanln()
	//using this to let the go routine finish

}

func printmsg(msg string) {

	fmt.Println(msg)
}
