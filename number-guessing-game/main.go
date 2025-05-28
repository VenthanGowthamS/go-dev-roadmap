package main

import (
	"fmt"
	"math/rand"
)

func main() {

	number := rand.Intn(100) + 1

	var guess int

	attempts := 0

	fmt.Println("welcome to number guessing game my dear buddy")

	fmt.Println("I have selected the number between 1 & 100 please guess it")

	for {

		fmt.Println("Im the GO master ! enter your guess human buddy !")

		fmt.Scan(&guess)

		attempts++

		if guess < number {
			fmt.Println("guess is lesser than i guessed try again with something higher input buddy")
		} else if guess > number {
			fmt.Println("Your guess values is higher than i have choose guess for some low number")
		} else {
			fmt.Printf("correct guess in attemts :%d", attempts)
			break
		}

	}
}
