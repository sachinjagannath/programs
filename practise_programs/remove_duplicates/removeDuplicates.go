package main

import "fmt"

func removeDuplicates(num []int) []int {
	var result []int
	printed := make(map[int]bool)
	for _, num := range num {
		if !printed[num] {
			result = append(result, num)
			printed[num] = true
		}
	}
	return result
}

func main() {
	num := []int{1, 2, 1, 3, 4, 2, 5, 3, 6, 4, 2, 3}

	remDup := removeDuplicates(num)

	fmt.Println(remDup)
}
