package calculator

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b float32) float32 {
	return a * b
}

// it can not used since it is in small case
func divide(a, b int16) int16 {
	return a / b

}

func Greetings() {
	fmt.Println("Welcome to Calculator Package!!! \n \n Learner Buddy")
}
