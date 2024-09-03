package stacks

// Pa moves the top element from stackB to the top of stackA.
func Pa(stackA, stackB *[]int) {
	if len(*stackB) > 0 { // Check if stackB is not empty
		*stackA = append([]int{(*stackB)[0]}, *stackA...) // Prepend the top element of stackB to stackA
		*stackB = (*stackB)[1:]                           // Remove the top element from stackB
	}
}

// Pb moves the top element from stackA to the top of stackB.
func Pb(stackA, stackB *[]int) {
	if len(*stackA) > 0 { // Check if stackA is not empty
		*stackB = append([]int{(*stackA)[0]}, *stackB...) // Prepend the top element of stackA to stackB
		*stackA = (*stackA)[1:]                           // Remove the top element from stackA
	}
}

// Sa swaps the first two elements of stackA.
func Sa(stackA *[]int) {
	if len(*stackA) > 1 { // Ensure there are at least two elements to swap
		(*stackA)[0], (*stackA)[1] = (*stackA)[1], (*stackA)[0] // Swap the first two elements
	}
}

// Sb swaps the first two elements of stackB.
func Sb(stackB *[]int) {
	if len(*stackB) > 1 { // Ensure there are at least two elements to swap
		(*stackB)[0], (*stackB)[1] = (*stackB)[1], (*stackB)[0] // Swap the first two elements
	}
}

// Ss swaps the first two elements of both stackA and stackB.
func Ss(stackA, stackB *[]int) {
	Sa(stackA) // Swap the first two elements of stackA
	Sb(stackB) // Swap the first two elements of stackB
}

// Ra rotates stackA upwards, moving the first element to the last position.
func Ra(stackA *[]int) {
	if len(*stackA) > 1 { // Ensure there is more than one element to rotate
		*stackA = append((*stackA)[1:], (*stackA)[0]) // Move the first element to the end of stackA
	}
}

// Rb rotates stackB upwards, moving the first element to the last position.
func Rb(stackB *[]int) {
	if len(*stackB) > 1 { // Ensure there is more than one element to rotate
		*stackB = append((*stackB)[1:], (*stackB)[0]) // Move the first element to the end of stackB
	}
}

// Rr rotates both stackA and stackB upwards.
func Rr(stackA, stackB *[]int) {
	Ra(stackA) // Rotate stackA upwards
	Rb(stackB) // Rotate stackB upwards
}

// Rra rotates stackA downwards, moving the last element to the first position.
func Rra(stackA *[]int) {
	if len(*stackA) > 1 { // Ensure there is more than one element to rotate
		*stackA = append([]int{(*stackA)[len(*stackA)-1]}, (*stackA)[:len(*stackA)-1]...) // Move the last element to the beginning of stackA
	}
}

// Rrb rotates stackB downwards, moving the last element to the first position.
func Rrb(stackB *[]int) {
	if len(*stackB) > 1 { // Ensure there is more than one element to rotate
		*stackB = append([]int{(*stackB)[len(*stackB)-1]}, (*stackB)[:len(*stackB)-1]...) // Move the last element to the beginning of stackB
	}
}

// Rrr rotates both stackA and stackB downwards.
func Rrr(stackA, stackB *[]int) {
	Rra(stackA) // Rotate stackA downwards
	Rrb(stackB) // Rotate stackB downwards
}
