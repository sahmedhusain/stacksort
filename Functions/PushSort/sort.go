package sort

import (
	instruction "PS/Functions/PushInstructions"
	stacks "PS/Functions/PushStacks"
)

func SortStack(stackA, stackB *[]int) []string {
	var instructions []string

	if len(*stackA) == 2 {
		for !IsSorted(*stackA) {
			stacks.Sa(stackA)
			instructions = append(instructions, "sa")
		}

	} else if len(*stackA) == 3 {
		minIndex := instruction.FindMinIndex(*stackA)
		maxIndex := instruction.FindMaxIndex(*stackA)

		if (minIndex == 0 && maxIndex == 1) || (minIndex == 2 && maxIndex == 0) {
			if minIndex == 0 && maxIndex == 1 {
				stacks.Sa(stackA)
				instructions = append(instructions, "sa")

				stacks.Ra(stackA)
				instructions = append(instructions, "ra")

			} else if minIndex == 2 && maxIndex == 0 {
				stacks.Ra(stackA)
				instructions = append(instructions, "ra")

				stacks.Sa(stackA)
				instructions = append(instructions, "sa")
			}

		} else {
			if minIndex == 1 && maxIndex == 0 {
				stacks.Ra(stackA)
				instructions = append(instructions, "ra")

			} else if minIndex == 1 && maxIndex == 2 {
				stacks.Sa(stackA)
				instructions = append(instructions, "sa")

			} else if minIndex == 2 && maxIndex == 1 {
				stacks.Rra(stackA)
				instructions = append(instructions, "rra")

			}
		}

	} else {
		for !IsSorted(*stackA) {
			minIndex := instruction.FindMinIndex(*stackA)

			if minIndex == 0 {
				stacks.Pb(stackA, stackB)
				instructions = append(instructions, "pb")
			} else if minIndex == 1 {
				stacks.Sa(stackA)
				instructions = append(instructions, "sa")
			} else if minIndex <= len(*stackA)/2 {
				stacks.Ra(stackA)
				instructions = append(instructions, "ra")
			} else {
				stacks.Rra(stackA)
				instructions = append(instructions, "rra")
			}
		}

	}

	for len(*stackB) > 0 {
		stacks.Pa(stackA, stackB)
		instructions = append(instructions, "pa")
	}

	return instructions
}

func IsRepeated(stack []int) bool {
	seen := make(map[int]bool)
	for _, num := range stack {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

func IsSorted(stack []int) bool {
	for i := 0; i < len(stack)-1; i++ {
		if stack[i] > stack[i+1] {
			return false
		}
	}
	return true
}
