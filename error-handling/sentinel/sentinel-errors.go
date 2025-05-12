package main

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

var ErrInvalidID = errors.New("id is invalid")

// handling sentinel errors
func getUser(id int) (string, error) {

	//we have to declare this as a Errors

	if id < 1 {
		return "", ErrInvalidID
	} else if id != 1 {
		return "", ErrUserNotFound
	} else {
		return "Mr.Coder", nil
	}

}

func main() {

	fmt.Println("we are going to learn sentinel errors")
	username, err := getUser(-2)

	if errors.Is(err, ErrUserNotFound) {
		fmt.Println("sentinel->user not found error handle it accordingly")
	} else if err != nil {
		fmt.Println("non sentinel error occured-->", err)

	} else {
		fmt.Println("user is -->", username)
	}

	username2, err2 := getUser(2)

	if errors.Is(err2, ErrUserNotFound) {
		fmt.Println("user not found error handle it accordingly")
	} else if err2 != nil {
		fmt.Println("")

	} else {
		fmt.Println("user is -->", username2)
	}

	username3, err3 := getUser(1)

	if errors.Is(err3, ErrUserNotFound) {
		fmt.Println("user not found error handle it accordingly")
	} else if err3 != nil {
		fmt.Printf("Someother error not defined as sentinel %v", err3)

	} else {
		fmt.Println("user is -->", username3)
	}

}
