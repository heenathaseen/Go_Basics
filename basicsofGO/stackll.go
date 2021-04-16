package main

import (
	"fmt"
)

type Stack struct {
	top *Element
	size int
}

type Element struct {
	value interface{} 
	next *Element
}


func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

func main() {
	stack := new(Stack)
	
	stack.Push("Things")
	stack.Push("and")
	stack.Push("Stuff")

	fmt.Println(stack.Len())

	
	for stack.Len() > 0 {
	
		fmt.Println(stack.Pop())

	}
	fmt.Println(stack.Len())
}