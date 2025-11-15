package main

import "fmt"

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

func squares(nums <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for num := range nums {
			out <- num * num
		}
		close(out)
	}()
	return out
}

func main() {
	nums := generate(1, 2, 3, 4)
	square := squares(nums)

	for v := range square {
		fmt.Println(v)
	}
}
