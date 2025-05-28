package main

import (
	"fmt"
)

func main() {

	//3*# is a size of the 2d array
	var board [3][3]string
	printboard(board)
}

func printboard(board [3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			//Since the board is initialized with strings and not yet filled with "X" or "O", the default is "".

			if board[i][j] == "" {
				fmt.Print("- ")
				//If the cell is empty (""), you print "- " to show a dash (indicating an empty slot).
			} else {
				fmt.Println(" ")
				//" " = space character (used to separate printed values)
			}
		}
		fmt.Println()
	}
}
