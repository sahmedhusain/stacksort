package instruction

// FindMinIndex returns the index of the smallest element in the stack
func FindMinIndex(stack []int) int {
	minIndex := 0                     // Assume the first element is the minimum
	for i := 1; i < len(stack); i++ { // Iterate over the stack starting from the second element
		if stack[i] < stack[minIndex] { // If the current element is smaller than the current minimum
			minIndex = i // Update minIndex to the current element's index
		}
	}
	return minIndex // Return the index of the smallest element
}

// FindMaxIndex returns the index of the largest element in the stack
func FindMaxIndex(stack []int) int {
	maxIndex := 0                     // Assume the first element is the maximum
	for i := 1; i < len(stack); i++ { // Iterate over the stack starting from the second element
		if stack[i] > stack[maxIndex] { // If the current element is larger than the current maximum
			maxIndex = i // Update maxIndex to the current element's index
		}
	}
	return maxIndex // Return the index of the largest element
}
