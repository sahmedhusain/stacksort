package main

import (
	sort "PS/Functions/PushSort"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	input := os.Args[1]

	elements := strings.Split(input, " ")

	var stackA []int

	for _, elem := range elements {
		num, err := strconv.Atoi(elem)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
		stackA = append(stackA, num)
	}

	var stackB []int

	if sort.IsRepeated(stackA) {
		fmt.Fprintln(os.Stderr, "Error")
		return

	} else if sort.IsSorted(stackA) {
		return

	}

	instructions := sort.SortStack(&stackA, &stackB)

	in := ""

	for _, instr := range instructions {
		in = in + instr + "\\n"
		fmt.Println(instr)
	}
}
