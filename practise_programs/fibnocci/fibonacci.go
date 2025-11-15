package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1

	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	n := 6
	num := fib(n)
	fmt.Println(num)
}
