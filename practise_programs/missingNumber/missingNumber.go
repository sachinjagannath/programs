package main

import "fmt"

func findMissingNumber(num []int) int {
	n := len(num)
	missing := n

	for i, num := range num {
		missing ^= i ^ num
	}
	return missing
}

func main() {
	num := []int{3, 0, 1}
	missing := findMissingNumber(num)
	fmt.Println(missing)
}
