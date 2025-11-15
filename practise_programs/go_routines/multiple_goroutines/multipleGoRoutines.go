package main

import (
	"fmt"
	"time"
)

func worker(i int) {
	fmt.Println(i)
}

func main() {
	for i := 0; i < 10; i++ {
		go worker(i)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("go routine from main")
}
