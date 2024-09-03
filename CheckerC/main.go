package main

import (
	chstack "PS/Functions/CheckerStacks"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	a, err := chstack.ParseArguments(strings.Split(os.Args[1], " "))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}
	b := &chstack.Stack{}

	var instructions []string
	for {
		var instruction string
		if _, err := fmt.Scanln(&instruction); err != nil {
			break
		}
		instructions = append(instructions, instruction)
	}

	for _, instruction := range instructions {
		if err := chstack.ExecuteInstruction(instruction, a, b); err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
	}

	if a.IsSorted() && len(b.Elements) == 0 {

		fmt.Println("OK")
	} else {

		fmt.Println("KO")
	}
}
