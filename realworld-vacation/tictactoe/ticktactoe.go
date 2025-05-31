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

	for i := 0; i < 9; i++ {

		fmt.Println("printing current board:")
		printboard(board)

		fmt.Println("Enter row (0-2)")
		fmt.Scan(&row)

		fmt.Println("enter coloumn (0-2)")
		fmt.Scan(&col)

		if row >= 0 && row <= 2 && col >= 0 && col <= 2 && board[row][col] == "" {
			//update the board with player
			board[row][col] = player

			if checkwinner(board, player) {
				printboard(board)
				fmt.Printf("player %s wins", player)
				return

			}
		} else {
			fmt.Println("invalid value ! please try again")
			i--
			continue
		}
		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}
	}
	fmt.Println("updated board:")
	printboard(board)
	fmt.Println("Final board:")
	printboard(board)
	fmt.Println("It's a draw!")
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

func checkwinner(board [3][3]string, player string) bool {

	//check coloumn
	for i := 0; i < 3; i++ {

		if board[0][i] == player && board[1][i] == player && board[2][i] == player {

			return true

		}

	}

	//check row
	for i := 0; i < 3; i++ {

		if board[i][0] == player && board[i][1] == player && board[i][2] == player {

			return true

		}

	}

	//check diagnol 1

	if board[0][0] == player && board[1][1] == player && board[2][2] == player {

		return true

	}
	//check diagnol 2
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {

		return true

	}

	return false

}
