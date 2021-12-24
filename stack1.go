package main

import (
	"errors"
	"fmt"
	"sync"
)

type Stack struct {
	s    []int
	lock sync.Mutex
}

func main() {
	var stack Stack
	stack.Push(1)
	ele, err := stack.Pop()
	if err != nil {

	}
	fmt.Println(ele)
}

func (s *Stack) Push(i int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, i)
}

func (s *Stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	length := len(s.s)
	if length == 0 {
		return 0, errors.New("Stack EMpty")
	}

	ele := s.s[length-1]
	s.s = s.s[:length-1]
	return ele, nil
}
