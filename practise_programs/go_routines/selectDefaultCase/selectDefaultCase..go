package main

import "fmt"

func main() {
	ch := make(chan int)

	select {
	case ch <- 1:
		fmt.Println("Semt 1")
	default:
		fmt.Println("channel busy, skip end")
	}

	select {
	case val := <-ch:
		fmt.Println("received", val)
	default:
		fmt.Println("Channel empty, skip receive")
	}
}
