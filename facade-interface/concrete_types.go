package main

//"This is the entry point for an executable application."

import (
	"fmt"
)

// In Go, methods for a struct are defined separately, after the struct definition.
// The (c *cpu) part is called the "receiver" and it tells Go that this function belongs to the cpu struct.
// The receiver comes before the function name and behaves like a reference to the struct instance (like 'this' in other languages).
// Using *cpu means we're using a pointer receiver, allowing us to modify the original struct's fields.
// The receiver (like a reference) is on the left side, and the method parameters go on the right side.

type cpu struct {
	status string
}

func (c *cpu) Freeze() {
	fmt.Println("Inside Freeze(): setting status to 'frozen'")
	c.status = "frozen"
}

func (c *cpu) Jump(position int) {
	fmt.Printf("\n CPU position is %d", position)
}

func (c *cpu) Execute() {
	fmt.Println("\n computer is executing sir")
}

type memory struct{}

func (m *memory) Load(position int, data string) {
	fmt.Printf("\nLoading in this data %s in %d position", data, position)
}

type harddisk struct{}

func (h *harddisk) Read(position int) string {
	data := "OS boot sector"
	fmt.Printf("\nReading at position %d", position)
	return data
}
