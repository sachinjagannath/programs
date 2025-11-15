package main

import "fmt"

type Stack struct {
	Items []int
}

func (s *Stack) Push(item int) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() int {
	if len(s.Items) == 0 {
		fmt.Println("Stack is empty...")
		return -1
	}
	top := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return top
}

func (s *Stack) Peek() (int, error) {
	if len(s.Items) == 0 {
		fmt.Println("Stack is empty...")
	}
	return s.Items[len(s.Items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}

func main() {
	var stack Stack
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Stack: ", stack.Items)

	top, _ := stack.Peek()
	fmt.Println("Peeked item is ", top)

	popped := stack.Pop()
	fmt.Println("Popped items is ", popped)

	fmt.Println("Final Items...")
	fmt.Println(stack.Items)

	fmt.Println("Checking if stack is empty or not")
	stack.IsEmpty()
}
