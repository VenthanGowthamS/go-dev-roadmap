package main

import "fmt"

func main() {

	fmt.Println("Basic for loop like (i=0,i<5,i++) ")

	for i := 0; i < 5; i++ {
		fmt.Println("values are i =", i)
	}

	fmt.Println("\n for loop as a while loop(condition only)")
	count := 0
	for count < 5 {
		fmt.Println("count is", count)
		count++
	}
	//adding count++ here instead the above will makeit infinite

	fmt.Println("\n for loop as infinite loop")
	num := 0
	for {
		if num > 3 {
			break

		}

		fmt.Println("number value =", num)
		num++
	}

}
