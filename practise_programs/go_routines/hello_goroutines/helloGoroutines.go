package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Hello from go routines...")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("exiting from go routines...")

}
