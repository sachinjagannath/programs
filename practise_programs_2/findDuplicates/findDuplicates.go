package main

import "fmt"

func findDuplicates(n []int) map[int]int {

	dupMap := make(map[int]int)
	for _, v := range n {
		dupMap[v]++
	}
	return dupMap
}

func main() {
	n := []int{1, 2, 4, 2, 1, 5, 6, 4}
	duplicates := findDuplicates(n)

	for i, v := range duplicates {
		if v > 1 {
			fmt.Printf("%d ", i)
		}
	}
}
