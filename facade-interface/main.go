package main // Declares this file belongs to the "main" package, which is the entry point of the Go program

func main() {
	pc := NewComputer() // Create a new Computer instance by calling the constructor function
	pc.StartComputer()  // Start the computer by invoking the method that operates the subsystems
}
