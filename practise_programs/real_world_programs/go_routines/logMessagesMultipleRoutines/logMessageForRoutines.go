package main

import (
	"fmt"
	"time"
)

func logger(ch <-chan string) {
	for msg := range ch {
		fmt.Println(time.Now().Format("15:05:05"), msg)
	}
}

func main() {
	logCh := make(chan string, 10)
	go logger(logCh)

	for i := 0; i <= 5; i++ {
		go func(id int) {
			logCh <- fmt.Sprintf("Worker %d started", id)
			time.Sleep(time.Second)
			logCh <- fmt.Sprintf("Worker %d done", id)
		}(i)
	}
	time.Sleep(2 * time.Second)
	close(logCh)
}
