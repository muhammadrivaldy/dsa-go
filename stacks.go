package main

import "fmt"

type stack struct {
	values []int
}

func (s *stack) print() {
	fmt.Printf("Output: %v\n", s.values)
}

func (s *stack) push(val int) {
	fmt.Println("Push:", val)
	s.values = append(s.values, val)
}

func (s *stack) pop() int {
	fmt.Println("Pop")
	if len(s.values) == 0 {
		return -1
	}

	val := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return val
}
