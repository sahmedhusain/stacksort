package main

import (
	sort "stacksort/Functions/PushSort" // Import the sorting and utility functions from the PushSort package
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Check if at least one argument is provided
	if len(os.Args) < 2 {
		return // Exit the program if no arguments are provided
	}

	input := os.Args[1] // Get the first argument passed to the program, expected to be a string of numbers

	// Split the input string by spaces to separate the individual numbers
	elements := strings.Split(input, " ")

	var stackA []int // Initialize an empty slice to hold the integers for stackA

	// Convert each element from string to integer and append it to stackA
	for _, elem := range elements {
		num, err := strconv.Atoi(elem) // Convert string to integer
		if err != nil {                // Check for conversion error
			fmt.Fprintln(os.Stderr, "Error") // Print error message to standard error
			return                           // Exit the program if there's an error
		}
		stackA = append(stackA, num) // Add the converted integer to stackA
	}

	var stackB []int // Initialize an empty slice to hold the integers for stackB

	// Check if the stack contains any repeated elements
	if sort.IsRepeated(stackA) {
		fmt.Fprintln(os.Stderr, "Error") // Print error message to standard error
		return                           // Exit the program if duplicates are found

		// Check if the stack is already sorted
	} else if sort.IsSorted(stackA) {
		return // Exit the program if the stack is already sorted
	}

	// Sort the stackA using the SortStack function, and store the instructions in a slice
	instructions := sort.SortStack(&stackA, &stackB)

	in := "" // Initialize an empty string to accumulate instructions

	// Iterate over the instructions and print each one
	for _, instr := range instructions {
		in = in + instr + "\\n" // Append the instruction to the accumulating string
		fmt.Println(instr)      // Print the instruction to standard output
	}
}
