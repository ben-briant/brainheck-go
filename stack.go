package main

import "fmt"

type LoopMarker struct {
	marker int
	next   *LoopMarker
}

type Stack struct {
	top *LoopMarker
}

func (s *Stack) Push(marker int) {
	newMarker := new(LoopMarker)
	newMarker.marker = marker
	newMarker.next = s.top
	s.top = newMarker
}

func (s *Stack) Pop() int {
	if s.top == nil {
		panic("No elements in the stack!")
	}
	poppedMarker := s.top
	s.top = s.top.next
	return poppedMarker.marker
}

func (s *Stack) Show() {
	mark := s.top
	for mark != nil {
		fmt.Printf("%d -> ", mark.marker)
		mark = mark.next
	}
	fmt.Println("X")
}

func StackNew() *Stack {
	return new(Stack)
}
