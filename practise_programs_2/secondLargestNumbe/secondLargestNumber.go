package main

import "math"

func secondLargestNumber(num []int) int {
	l, sl := math.MinInt, math.MinInt

	for _, n := range num {
		if n > l {
			l = n
			sl = l
		}
	}

}

func main() {
	num := []int{1, 2, 3, 4, 2, 7, 6, 8, 23, 234, 6}

}
