package main

import (
	chstack "stacksort/Functions/CheckerStacks" // Import the custom package for stack operations, aliasing it as 'chstack'
	"fmt"                                // Import the fmt package for formatted I/O operations
	"os"                                 // Import the os package to handle command-line arguments and standard input/output
	"strings"                            // Import the strings package for string manipulation
)

func main() {
	// Check if the program was run with at least one argument
	if len(os.Args) < 2 {
		return // Exit the program if no argument is provided
	}

	// Parse the first command-line argument as a space-separated list of integers
	a, err := chstack.ParseArguments(strings.Split(os.Args[1], " "))
	if err != nil { // If there is an error in parsing the arguments
		fmt.Fprintln(os.Stderr, "Error") // Print "Error" to the standard error stream
		return                           // Exit the program
	}

	// Initialize an empty stack 'b' to use as the secondary stack in operations
	b := &chstack.Stack{}

	// Prepare to collect a list of instructions from user input
	var instructions []string
	for {
		var instruction string
		// Read an instruction from standard input
		if _, err := fmt.Scanln(&instruction); err != nil {
			break // Exit the loop if input ends (EOF or error)
		}
		// Add the instruction to the list of instructions
		instructions = append(instructions, instruction)
	}

	// Execute each instruction in the collected list
	for _, instruction := range instructions {
		// Apply the instruction on stacks 'a' and 'b'
		if err := chstack.ExecuteInstruction(instruction, a, b); err != nil {
			fmt.Fprintln(os.Stderr, "Error") // Print "Error" to standard error if instruction execution fails
			return                           // Exit the program
		}
	}

	// Check if stack 'a' is sorted and stack 'b' is empty
	if a.IsSorted() && len(b.Elements) == 0 {
		fmt.Println("OK") // Print "OK" if the conditions are met, meaning the sequence is correct
	} else {
		fmt.Println("KO") // Print "KO" if the sequence is incorrect
	}
}
