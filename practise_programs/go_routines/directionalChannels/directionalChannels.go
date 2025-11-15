package main

import "fmt"

func send(ch chan<- int) {
	for i := 0; i <= 3; i++ {
		ch <- i
	}
	close(ch)
}

func receive(ch <-chan int) {
	for c := range ch {
		fmt.Println(c)
	}
}

func main() {
	ch := make(chan int)
	send(ch)
	receive(ch)

}
