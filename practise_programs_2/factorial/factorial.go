package main

import "fmt"

func getFactorial(num int) int {
	n := 1

	for num > 0 {
		n = n * num
		num--
	}
	return n
}

func main() {
	var num int
	fmt.Println("Enter a number to get the factorial of.. ")
	fmt.Scan(&num)

	fact := getFactorial(num)
	fmt.Printf("Factorial of a number %d is %d", num, fact)
}
