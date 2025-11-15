package main

import (
	"fmt"
	"sync"
)

func worker(id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	sem <- struct{}{}
	fmt.Println("worker", id)
	<-sem
}

func main() {
	const maxConcurrent = 2
	sem := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, sem, &wg)
	}
	wg.Wait()
}
