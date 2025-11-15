package main

import "fmt"

func main() {
	a, b := 0, 1
	n := 2
	for {
		if n > 9 {
			break
		}
		fmt.Printf("%d ", a)
		a, b = b, a+b
		n++
	}
	fmt.Println()
}
