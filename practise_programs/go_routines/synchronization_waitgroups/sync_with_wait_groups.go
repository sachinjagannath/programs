package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Printf("\nWorker %d started", i)
	time.Sleep(time.Second)
	fmt.Printf("\nWorker %d done", i)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}
	wg.Wait()
	fmt.Println("\nAll workers finished")
}
