package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg)
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	msg := "printing message from print method"
	go printMessage(msg)
	time.Sleep(8000 * time.Millisecond)
	fmt.Println("go routine from main")
}
