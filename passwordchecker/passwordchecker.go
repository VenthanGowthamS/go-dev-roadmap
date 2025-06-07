package main

import (
	"fmt"
	"unicode"
)

func main() {

	var pw string

	fmt.Println("enter the password")

	fmt.Scanln(&pw)

	if passwordchecker(pw) {
		fmt.Println("password is strong")
	} else {
		fmt.Println("password is weak")
	}

}

func passwordchecker(pw string) bool {

	var hasalpha, hasUpper, hasSpecial, hasnumber bool

	for _, char := range pw {
		switch {

		case unicode.IsUpper(char):

			hasUpper = true
		case unicode.IsDigit(char):

			hasnumber = true

		case unicode.IsLetter(char):

			hasalpha = true

		case unicode.IsPunct(char) || unicode.IsSymbol(char):

			hasSpecial = true

		}

	}

	return len(pw) > 8 && hasalpha && hasUpper && hasnumber && hasSpecial
}
