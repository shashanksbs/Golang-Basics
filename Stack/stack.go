package main

import "fmt"

type Stack []string

//push -
//pop
//dispaly
func main() {
	var stack Stack
	stack.push("sbs")
	element, bln := stack.Pop()
	if bln {
		fmt.Println("element", element)
	}
}

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(str string) {

	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
