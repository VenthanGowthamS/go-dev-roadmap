package main

import (
	"fmt"
)

func main() {

	//3*# is a size of the 2d array
	var board [3][3]string

	//var row int
	//var col int equals
	var row, col int
	player := "X"
	fmt.Println("printing current board:")
	printboard(board)

	fmt.Println("Enter row (0-2)")
	fmt.Scan(&row)

	fmt.Println("enter coloumn (0-2)")
	fmt.Scan(&col)

	if row >= 0 && row <= 2 && col >= 0 && col <= 2 && board[row][col] == "" {
		//update the board wiht player
		board[row][col] = player
	} else {
		fmt.Println("invalid value ! please try again")
	}
	fmt.Println("updated board:")
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
				fmt.Print(board[i][j] + " ")
				//" " = space character (used to separate printed values)
			}
		}
		fmt.Println()
	}
}
