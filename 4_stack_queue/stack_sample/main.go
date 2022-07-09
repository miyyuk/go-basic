package main

import "fmt"

var max int = 100000

type Stack struct {
	data []int
}

type StackI interface {
	isEmpty() bool
	isFull() bool
	push(int)
	pop() int
}

var _ StackI = &Stack{}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) isFull() bool {
	return len(s.data) == max
}

func (s *Stack) push(n int) {
	if s.isFull() {
		fmt.Println("error: stack is full")
		return
	}
	s.data = append(s.data, n)
}

func (s *Stack) pop() int {
	if s.isEmpty() {
		fmt.Println("error: stack is empty")
		return 0
	}
	last := len(s.data) - 1
	popData := s.data[last]
	s.data = s.data[:last]
	return popData
}

func main() {
	var s *Stack = &Stack{} // Stackを初期化

	s.push(3) // [] -> [3]
	s.push(5) // [3] -> [3, 5]
	s.push(7) // [3, 5] -> [3, 5, 7]

	fmt.Println(s.pop()) // [3, 5, 7] -> [3, 5] で 7 を出力
	fmt.Println(s.pop()) // [3, 5] -> [3] で 5 を出力

	s.push(9)  // [3] -> [3, 9]
	s.push(11) // [3, 9] -> [3, 9, 11]

	s.pop() // [3, 9, 11] -> [3, 9]
	s.pop() // [3, 9] -> [3]
	s.pop() // [3] -> []

	if s.isEmpty() {
		fmt.Println("empty")
	} else {
		fmt.Println("not empty")
	}
}
