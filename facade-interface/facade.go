/*
This is a simplified simulation of how a computer boots up:

CPU freezes current state

Memory loads the boot instructions

CPU jumps to a memory location

CPU starts executing instructions

Hard Drive provides the data (e.g., boot sector)

The goal here is not to build a real OS,
but to illustrate how complex subsystems (CPU, Memory, Hard Drive) can be controlled through a simplified interface
— the Facade.*/

package main

import (
	"fmt"
)

// Computer is a high-level struct that uses CPU, Memory, and Harddisk interfaces.
type Computer struct {
	cpu      CPU      // Interface type for CPU component
	memory   Memory   // Interface type for Memory component
	harddisk Harddisk // Interface type for Harddisk component
}

// NewComputer is a constructor function that returns a pointer to a new Computer object.
func NewComputer() *Computer {
	c := &cpu{status: "idle"} // Create a pointer to a concrete cpu struct and set its initial status to "idle"

	// Print the address and current status of the cpu object before assigning it to the interface
	fmt.Printf("Before assigning to interface: cpu address = %p, status = %s\n", c, c.status)

	// Return a new Computer struct initialized with the components
	return &Computer{
		cpu:      c,           // Assign the cpu pointer to the CPU interface field
		memory:   &memory{},   // Create a pointer to the concrete memory struct and assign it
		harddisk: &harddisk{}, // Create a pointer to the concrete harddisk struct and assign it
	}
}

// StartComputer is a method on the Computer struct that starts the computer using its components
func (comp *Computer) StartComputer() {
	comp.cpu.Freeze()             // Step 1: Freeze the CPU
	data := comp.harddisk.Read(0) // Step 2: Read data from position 0 of the harddisk
	comp.memory.Load(0, data)     // Step 3: Load the data into memory at position 0
	comp.cpu.Jump(0)              // Step 4: Make the CPU jump to position 0
	comp.cpu.Execute()            // Step 5: Execute the instructions

	// Type assertion to retrieve the underlying concrete *cpu from the CPU interface
	if realCPU, ok := comp.cpu.(*cpu); ok {
		// If successful, print the memory address and status after starting
		fmt.Printf("After start: cpu address = %p, status = %s\n", realCPU, realCPU.status)
	}
}
