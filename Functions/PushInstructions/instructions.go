package instruction

func FindMinIndex(stack []int) int {
	minIndex := 0
	for i := 1; i < len(stack); i++ {
		if stack[i] < stack[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func FindMaxIndex(stack []int) int {
	maxIndex := 0
	for i := 1; i < len(stack); i++ {
		if stack[i] > stack[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}