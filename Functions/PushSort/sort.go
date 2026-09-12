package sort

import (
	instruction "stacksort/Functions/PushInstructions" // Import the package for instruction-related functions
	stacks "stacksort/Functions/PushStacks"            // Import the package for stack operations
)

// SortStack sorts the stackA using an auxiliary stackB and returns the list of instructions used
func SortStack(stackA, stackB *[]int) []string {
	var instructions []string // To store the sequence of operations

	// Case for sorting a stack with 2 elements
	if len(*stackA) == 2 {
		for !IsSorted(*stackA) { // If stackA is not sorted
			stacks.Sa(stackA)                         // Swap the two elements
			instructions = append(instructions, "sa") // Record the operation
		}

		// Case for sorting a stack with 3 elements
	} else if len(*stackA) == 3 {
		minIndex := instruction.FindMinIndex(*stackA) // Find the index of the smallest element
		maxIndex := instruction.FindMaxIndex(*stackA) // Find the index of the largest element

		// Check specific scenarios for the arrangement of min and max elements
		if (minIndex == 0 && maxIndex == 1) || (minIndex == 2 && maxIndex == 0) {
			// Case 1: Minimum at index 0 and Maximum at index 1
			if minIndex == 0 && maxIndex == 1 {
				stacks.Sa(stackA) // Swap to correct the order
				instructions = append(instructions, "sa")

				stacks.Ra(stackA) // Rotate to move the minimum element to the correct position
				instructions = append(instructions, "ra")

				// Case 2: Minimum at index 2 and Maximum at index 0
			} else if minIndex == 2 && maxIndex == 0 {
				stacks.Ra(stackA) // Rotate to move the minimum element to the correct position
				instructions = append(instructions, "ra")

				stacks.Sa(stackA) // Swap to correct the order
				instructions = append(instructions, "sa")
			}

			// Handle other scenarios for the arrangement of min and max elements
		} else {
			if minIndex == 1 && maxIndex == 0 {
				stacks.Ra(stackA) // Rotate stackA to bring the minimum to the correct position
				instructions = append(instructions, "ra")

			} else if minIndex == 1 && maxIndex == 2 {
				stacks.Sa(stackA) // Swap the top two elements
				instructions = append(instructions, "sa")

			} else if minIndex == 2 && maxIndex == 1 {
				stacks.Rra(stackA) // Reverse rotate to move the minimum to the correct position
				instructions = append(instructions, "rra")

			}
		}

		// Case for sorting a stack with more than 3 elements
	} else {
		for !IsSorted(*stackA) { // Continue until stackA is sorted
			minIndex := instruction.FindMinIndex(*stackA) // Find the index of the smallest element

			if minIndex == 0 { // If the smallest element is at the top
				stacks.Pb(stackA, stackB) // Push it to stackB
				instructions = append(instructions, "pb")
			} else if minIndex == 1 { // If the smallest element is in the second position
				stacks.Sa(stackA) // Swap to move it to the top
				instructions = append(instructions, "sa")
			} else if minIndex <= len(*stackA)/2 { // If the smallest element is in the first half
				stacks.Ra(stackA) // Rotate stackA to bring the smallest to the top
				instructions = append(instructions, "ra")
			} else { // If the smallest element is in the second half
				stacks.Rra(stackA) // Reverse rotate stackA to bring the smallest to the top
				instructions = append(instructions, "rra")
			}
		}
	}

	// After sorting, push all elements from stackB back to stackA
	for len(*stackB) > 0 {
		stacks.Pa(stackA, stackB) // Push elements back from stackB to stackA
		instructions = append(instructions, "pa")
	}

	return instructions // Return the sequence of instructions performed
}

// IsRepeated checks if there are any duplicate elements in the stack
func IsRepeated(stack []int) bool {
	seen := make(map[int]bool)  // Create a map to track seen elements
	for _, num := range stack { // Iterate through the stack
		if seen[num] { // If the number is already seen
			return true // Return true, indicating a duplicate
		}
		seen[num] = true // Mark the number as seen
	}
	return false // Return false if no duplicates are found
}

// IsSorted checks if the stack is sorted in ascending order
func IsSorted(stack []int) bool {
	for i := 0; i < len(stack)-1; i++ { // Iterate through the stack
		if stack[i] > stack[i+1] { // If any element is greater than the next
			return false // Return false, indicating the stack is not sorted
		}
	}
	return true // Return true if the stack is sorted
}
