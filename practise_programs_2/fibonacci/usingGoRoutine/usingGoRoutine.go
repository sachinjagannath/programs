package main

import "fmt"

func printFib(ch chan<- int, num int) {

	a, b := 0, 1
	n := 2
	for n < num {
		ch <- a
		a, b = b, b+a
		n++
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go printFib(ch, 10)

	for c := range ch {
		fmt.Printf("%d ", c)
	}
	fmt.Println()
}
