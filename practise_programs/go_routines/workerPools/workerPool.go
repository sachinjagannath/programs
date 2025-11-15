package main

import (
	"fmt"
	"sync"
)

//cont (
//	noOfWorkers:=3
//	)

func worker(w int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d \n", w, j)
	}
}

func main() {
	noOfWorkers := 3
	jobs := make(chan int)
	var wg sync.WaitGroup

	for w := 1; w <= noOfWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
}
