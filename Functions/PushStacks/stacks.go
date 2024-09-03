package stacks

func Pa(stackA, stackB *[]int) {
	if len(*stackB) > 0 {
		*stackA = append([]int{(*stackB)[0]}, *stackA...)
		*stackB = (*stackB)[1:]
	}
}

func Pb(stackA, stackB *[]int) {
	if len(*stackA) > 0 {
		*stackB = append([]int{(*stackA)[0]}, *stackB...)
		*stackA = (*stackA)[1:]
	}
}

func Sa(stackA *[]int) {
	if len(*stackA) > 1 {
		(*stackA)[0], (*stackA)[1] = (*stackA)[1], (*stackA)[0]
	}
}

func Sb(stackB *[]int) {
	if len(*stackB) > 1 {
		(*stackB)[0], (*stackB)[1] = (*stackB)[1], (*stackB)[0]
	}
}

func Ss(stackA, stackB *[]int) {
	Sa(stackA)
	Sb(stackB)
}

func Ra(stackA *[]int) {
	if len(*stackA) > 1 {
		*stackA = append((*stackA)[1:], (*stackA)[0])
	}
}

func Rb(stackB *[]int) {
	if len(*stackB) > 1 {
		*stackB = append((*stackB)[1:], (*stackB)[0])
	}
}

func Rr(stackA, stackB *[]int) {
	Ra(stackA)
	Rb(stackB)
}

func Rra(stackA *[]int) {
	if len(*stackA) > 1 {
		*stackA = append([]int{(*stackA)[len(*stackA)-1]}, (*stackA)[:len(*stackA)-1]...)
	}
}

func Rrb(stackB *[]int) {
	if len(*stackB) > 1 {
		*stackB = append([]int{(*stackB)[len(*stackB)-1]}, (*stackB)[:len(*stackB)-1]...)
	}
}

func Rrr(stackA, stackB *[]int) {
	Rra(stackA)
	Rrb(stackB)
}
