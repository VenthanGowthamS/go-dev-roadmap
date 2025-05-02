package main

import "fmt"

//function called byitself to print
//0 1 1 2 3 ...
//Each number is the sum of the two before it:

func main() {

	fmt.Println("Recursive fibonnaci series")

	for i := 0; i < 10; i++ {
		println(fibonnaci(i))
	}
}

func fibonnaci(n int) int {

	if n <= 1 {
		return n
	}
	return fibonnaci(n-1) + fibonnaci(n-2)
}
