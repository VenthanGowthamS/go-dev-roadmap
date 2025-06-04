package main

import "fmt"

func main() {
	var number int

	fmt.Println("enter a number")
	fmt.Scanln(&number)

	if isprime(number) {

		fmt.Println("it is a prime number", number)

	} else {
		fmt.Println("is is not a prime number", number)
	}
}

func isprime(n int) bool {

	if n <= 1 {
		return false
	}

	for i := 2; i*1 < n; i++ {

		if n%i == 0 {
			return false

		}

	}
	return true
}
