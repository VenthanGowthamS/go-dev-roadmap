package main

import (
	"fmt"
	"os"
)

func main() {

	//file, err := os.Create("testfile<>.txt")
	file, err := os.Create("testfile.txt")
	if err != nil {
		fmt.Println("errro creating file", err)
		return
		//If return is used in the main function, it stops the program without needing to explicitly return a value.
		//Although the main function doesn't return anything, the return statement in the main function still serves an important purpose:
		//In Go, the main function does not have an explicit return type. Unlike other functions that might return values (e.g., int, string, error), the main function is special. It's the entry point of a Go program, and its signature is always:

	}
	//This program may leak resources because file.Close() is not called.
	// defer ensures the file is closed even if an error or return happens
	defer file.Close()

	fmt.Fprintf(file, "Hello Mr.coder")

}
