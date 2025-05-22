package main

import (
	"fmt"
	"sync"
)

//Use two goroutines to print even and odd numbers from 1 to 10.

func main() {

	var wg sync.WaitGroup
	//create wait group object
	wg.Add(2)
	//create counter with 2

	go printeven(&wg)
	//passing pointer to the function in the go routine

	//giving the memory address of &wg
	go printodd(&wg)
	wg.Wait()
	//wait till the go routine complete or the counter to 0
}

//since shared work group counter needs to be updated the

// * sync.waitgroup means im receiving a pointer to sync waitgroup
func printeven(wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 2; i < 10; i += 2 {

		fmt.Println("Even", i)

	}

}

func printodd(wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 1; i < 10; i += 2 {

		fmt.Println("odd", i)

	}
}
