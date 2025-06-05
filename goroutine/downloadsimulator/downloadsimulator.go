package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	wg := sync.WaitGroup{}

	numberOfDownloader := 5
	for i := 1; i <= numberOfDownloader; i++ {

		wg.Add(1)
		go downloadfile(i, &wg)
	}

	wg.Wait()
}

func downloadfile(id int, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Printf("%d is starting \n", id)
	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Nanosecond)
	fmt.Printf("%d is complete \n", id)

}
