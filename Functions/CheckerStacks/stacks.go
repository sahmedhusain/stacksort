package chstack

import (
	"fmt"     // Import the fmt package for formatted I/O operations
	"strconv" // Import the strconv package for converting strings to integers
)

// Stack is a struct that represents a stack of integers
type Stack struct {
	Elements []int // Elements holds the stack's values as a slice of integers
}

// Push adds an integer to the top of the stack (front of the slice)
func (s *Stack) Push(value int) {
	s.Elements = append([]int{value}, s.Elements...) // Prepend the value to the slice
}

// PUSH adds an integer to the bottom of the stack (end of the slice)
func (s *Stack) PUSH(value int) {
	s.Elements = append(s.Elements, value) // Append the value to the slice
}

// Pop removes and returns the top value of the stack
func (s *Stack) Pop() int {
	if len(s.Elements) == 0 { // Check if the stack is empty
		return -1 // Return -1 if the stack is empty
	}
	value := s.Elements[0]      // Get the top value (first element)
	s.Elements = s.Elements[1:] // Remove the top value from the stack
	return value                // Return the popped value
}

// Swap exchanges the top two elements of the stack
func (s *Stack) Swap() {
	if len(s.Elements) < 2 { // Ensure there are at least two elements to swap
		return // Do nothing if there are fewer than two elements
	}
	s.Elements[0], s.Elements[1] = s.Elements[1], s.Elements[0] // Swap the first two elements
}

// Rotate moves the top element of the stack to the bottom
func (s *Stack) Rotate() {
	if len(s.Elements) == 0 { // Check if the stack is empty
		return // Do nothing if the stack is empty
	}
	value := s.Elements[0]                // Get the top value
	copy(s.Elements, s.Elements[1:])      // Shift all elements to the left
	s.Elements[len(s.Elements)-1] = value // Place the top value at the bottom
}

// ReverseRotate moves the bottom element of the stack to the top
func (s *Stack) ReverseRotate() {
	if len(s.Elements) == 0 { // Check if the stack is empty
		return // Do nothing if the stack is empty
	}
	value := s.Elements[len(s.Elements)-1]                               // Get the bottom value
	s.Elements = append([]int{value}, s.Elements[:len(s.Elements)-1]...) // Move the bottom value to the top
}

// ParseArguments converts a list of strings into a Stack of integers
func ParseArguments(args []string) (*Stack, error) {
	a := &Stack{}              // Initialize an empty stack
	for _, arg := range args { // Iterate over each argument
		value, err := strconv.Atoi(arg) // Convert the argument to an integer
		if err != nil {                 // If conversion fails, return an error
			return nil, fmt.Errorf("error: %v", err)
		}
		a.PUSH(value) // Push the converted value onto the stack
	}
	return a, nil // Return the populated stack
}

// ExecuteInstruction performs the given instruction on stacks 'a' and 'b'
func ExecuteInstruction(instruction string, a, b *Stack) error {
	switch instruction { // Handle each instruction case
	case "sa":
		a.Swap() // Swap the top two elements of stack 'a'
	case "sb":
		b.Swap() // Swap the top two elements of stack 'b'
	case "ss":
		a.Swap() // Swap both 'a' and 'b'
		b.Swap()
	case "pa":
		if len(b.Elements) == 0 { // Ensure 'b' is not empty
			return fmt.Errorf("error: cannot pop from empty stack b")
		}
		a.Push(b.Pop()) // Pop from 'b' and push onto 'a'
	case "pb":
		if len(a.Elements) == 0 { // Ensure 'a' is not empty
			return fmt.Errorf("error: cannot pop from empty stack a")
		}
		b.Push(a.Pop()) // Pop from 'a' and push onto 'b'
	case "ra":
		a.Rotate() // Rotate stack 'a'
	case "rb":
		b.Rotate() // Rotate stack 'b'
	case "rr":
		a.Rotate() // Rotate both 'a' and 'b'
		b.Rotate()
	case "rra":
		a.ReverseRotate() // Reverse rotate stack 'a'
	case "rrb":
		b.ReverseRotate() // Reverse rotate stack 'b'
	case "rrr":
		a.ReverseRotate() // Reverse rotate both 'a' and 'b'
		b.ReverseRotate()
	default:
		return fmt.Errorf("error: invalid instruction %s", instruction) // Handle invalid instructions
	}
	return nil // Return nil if execution is successful
}

// IsSorted checks if the stack is sorted in ascending order
func (s *Stack) IsSorted() bool {
	for i := 0; i < len(s.Elements)-1; i++ { // Iterate through the stack
		if s.Elements[i] > s.Elements[i+1] { // Check if any element is greater than the next
			return false // Return false if the stack is not sorted
		}
	}
	return true // Return true if the stack is sorted
}
