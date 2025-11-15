package main

import (
	"fmt"
	"time"
)

func printNum(num int) {
	time.Sleep(time.Duration(num) * 2000 * time.Millisecond)
	fmt.Println(num)
}

func main() {
	for i := 0; i <= 3; i++ {
		go printNum(i)
	}
	time.Sleep(10500 * time.Millisecond)
}
