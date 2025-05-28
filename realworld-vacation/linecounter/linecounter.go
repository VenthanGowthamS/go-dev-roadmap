package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("testfile.txt")
	if err != nil {
		fmt.Println("error opneing file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	linecount := 0

	//while loop type for loop go
	for scanner.Scan() {
		linecount++
	}

	//✅ If the loop ended normally (EOF), scanner.Err() returns nil → nothing happens.
	// If the loop ended because of a read error, scanner.Err() returns a non-nil error → this block runs.

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file", err)
		return
	}

	fmt.Println("line count", linecount)

}
