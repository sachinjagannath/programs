package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "1"
	ch <- "2"

	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
