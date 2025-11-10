package main

import (
	"fmt"
	"net/http"
	"time"
)

func worker(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("error %w", err)
		return
	}
	err = resp.Body.Close()
	if err != nil {
		return
	}
	elapsed := time.Since(start)
	ch <- fmt.Sprintf("%s took %v", url, elapsed)
}

func main() {
	start := time.Now()
	ch := make(chan string)

	urls := []string{
		"https://golang.org", "https://example.com", "https://httpbin.org/get",
	}

	for _, url := range urls {
		go worker(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Println("Total Time: ", time.Since(start))

}
