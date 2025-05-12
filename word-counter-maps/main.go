package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	Scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the sentence")

	Scanner.Scan()

	input := Scanner.Text()

	//get user input and scan it
	lower := strings.ToLower(input)

	//making to to lower case

	words := strings.Fields(lower)

	//Storing as String in a slice

	//daclare a map

	wordcount := make(map[string]int)

	//get word count of the entered sentence by looping all the words buddy

	//word directly declared in the loop
	//The _ ignores the index (we don’t need it here).

	for _, word := range words {
		wordcount[word]++
	}

	//finaly venthan print the sentence enter with the values
	fmt.Println("\n words frequency count in the sentence you entered Mr.Coder")

	for word, count := range wordcount {

		fmt.Printf(" %s <---> %d", word, count)
		fmt.Println()
	}

}
