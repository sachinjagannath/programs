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
		fmt.Println("Stack is empty")
	}

	top := len(s.Items)
	s.Items = s.Items[:len(s.Items)-1]
	return top
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}

type Queue struct {
	stackIn  Stack
	stackOut Stack
}

func (q *Queue) Enqueue(item int) {
	q.stackIn.Push(item)
}

func (q *Queue) Dequeue() int {
	if q.stackOut.IsEmpty() {
		for !q.stackIn.IsEmpty() {
			q.stackOut.Push(q.stackIn.Pop())
		}
	}

	if q.stackOut.IsEmpty() {
		fmt.Println("Queue is empty")
		return -1
	}

	return q.stackOut.Pop()
}

func (q *Queue) Peek() int {
	if q.stackOut.IsEmpty() {
		for !q.stackIn.IsEmpty() {
			q.stackOut.Push(q.stackIn.Pop())
		}
	}
	if q.stackOut.IsEmpty() {
		fmt.Println("Queue is empty")
		return -1
	}

	return q.stackOut.Items[len(q.stackOut.Items)-1]
}
func main() {
	q := Queue{}

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	fmt.Println(q.Dequeue()) // 10
	q.Enqueue(40)
	fmt.Println(q.Dequeue()) // 20
	fmt.Println(q.Peek())    // 30
	fmt.Println(q.Dequeue()) // 30
	fmt.Println(q.Dequeue()) // 40
	fmt.Println(q.Dequeue())
}
