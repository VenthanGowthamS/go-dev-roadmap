package main

import (
	"bufio"
	"fmt"
	"os"
)

// learning switch case
// slice of the strings
func main() {

	//declare slice of the string name tasks ----here slice are nothing but list like in java & python

	var tasks []string

	//for standard output and input

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n ...This is to list function")
		fmt.Println("1. for adding a task")
		fmt.Println("2. for viewing a task")
		fmt.Println("3. for exit a task")
		fmt.Println("\n Enter your choice please")

		scanner.Scan()           //waiting for user input
		choice := scanner.Text() //user pressed enter with value
		switch choice {

		//now user choosen case we are going in side the each choice to do actions on

		case "1":

			fmt.Println("\n Enter the task you want to add")

			scanner.Scan()
			task := scanner.Text()
			tasks = append(tasks, task)
			fmt.Println("\n Congrats Venthan ! your choice of task added")

		case "2":

			fmt.Println("Im going to view the task in the list Venthan")

			//before just scanning for lenght and throw eero before

			if len(tasks) == 0 {
				fmt.Println("Oops Programmer ! No task to show")
			} else {
				//no scanner here just iterate the loop and print it

				for i, task := range tasks {

					fmt.Println("tasks are --->", task)
					i++
				}
				fmt.Println("hope you have seen your task! for each loop with range in go learned as well")
			}

		case "3":

			fmt.Println("existing from the choice selection")

			return

		default:
			fmt.Println("Choice not in the list. Please try again Mr.")
		}
	}

}
