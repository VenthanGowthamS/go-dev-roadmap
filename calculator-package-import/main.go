package main

import (
	"calculator-package-import/calculator"
	"fmt"
)

func main() {

	calculator.Greetings()

	addition := calculator.Add(45, 345345)
	subtraction := calculator.Subtract(3534534, 0)
	mutliplication := calculator.Multiply(5645.25, 54.02)

	fmt.Println("add value is >>>", addition)
	fmt.Println("subtract value is >>>", subtraction)
	fmt.Println("multiplied value is >>>", mutliplication)
}
