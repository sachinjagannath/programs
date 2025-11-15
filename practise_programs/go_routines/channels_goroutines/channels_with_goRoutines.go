package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello Go Routine"
	}()

	msg := <-ch
	fmt.Println(msg)
}
