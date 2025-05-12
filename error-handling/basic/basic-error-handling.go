package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("error handling when invalid values supplied to function")

	result1, err1 := divide(33, 1)

	if err1 == nil {
		fmt.Println("Value after division Mr.Coder>>>>>", result1)
	} else {
		fmt.Println("error returned-->", err1)
	}

	result2, err2 := divide(33, 0)

	if err2 == nil {
		fmt.Println("Value after division Mr.Coder>>>>>", result2)
	} else {
		fmt.Println("error returned-->", err2)
	}

}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("can not divide by 0")
	}
	return a / b, nil
}
