package chstack

import (
	"fmt"
	"strconv"
)

type Stack struct {
	Elements []int
}

func (s *Stack) Push(value int) {
	s.Elements = append([]int{value}, s.Elements...)
}
func (s *Stack) PUSH(value int) {
	s.Elements = append(s.Elements, value)
}

func (s *Stack) Pop() int {
	if len(s.Elements) == 0 {
		return -1
	}
	value := s.Elements[0]
	s.Elements = s.Elements[1:]
	return value
}

func (s *Stack) Swap() {
	if len(s.Elements) < 2 {
		return
	}
	s.Elements[0], s.Elements[1] = s.Elements[1], s.Elements[0]
}

func (s *Stack) Rotate() {
	if len(s.Elements) == 0 {
		return
	}
	value := s.Elements[0]
	copy(s.Elements, s.Elements[1:])
	s.Elements[len(s.Elements)-1] = value
}

func (s *Stack) ReverseRotate() {
	if len(s.Elements) == 0 {
		return
	}
	value := s.Elements[len(s.Elements)-1]
	s.Elements = append([]int{value}, s.Elements[:len(s.Elements)-1]...)
}

func ParseArguments(args []string) (*Stack, error) {
	a := &Stack{}
	for _, arg := range args {
		value, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("error: %v", err)
		}
		a.PUSH(value)
	}
	return a, nil
}

func ExecuteInstruction(instruction string, a, b *Stack) error {
	switch instruction {
	case "sa":
		a.Swap()

	case "sb":
		b.Swap()

	case "ss":
		a.Swap()
		b.Swap()

	case "pa":
		if len(b.Elements) == 0 {
			return fmt.Errorf("error: cannot pop from empty stack b")
		}
		a.Push(b.Pop())

	case "pb":
		if len(a.Elements) == 0 {
			return fmt.Errorf("error: cannot pop from empty stack a")
		}
		b.Push(a.Pop())

	case "ra":
		a.Rotate()

	case "rb":
		b.Rotate()

	case "rr":
		a.Rotate()
		b.Rotate()

	case "rra":
		a.ReverseRotate()

	case "rrb":
		b.ReverseRotate()

	case "rrr":
		a.ReverseRotate()
		b.ReverseRotate()

	default:
		return fmt.Errorf("error: invalid instruction %s", instruction)
	}
	return nil
}

func (s *Stack) IsSorted() bool {
	for i := 0; i < len(s.Elements)-1; i++ {
		if s.Elements[i] > s.Elements[i+1] {
			return false
		}
	}
	return true
}
