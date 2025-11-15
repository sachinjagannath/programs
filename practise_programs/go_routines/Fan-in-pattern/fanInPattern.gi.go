package main

import "fmt"

func generator(id int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- id*10 + i
		}
		close(ch)
	}()
	return ch
}

func fanIn(ch1, ch2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for ch1 != nil || ch2 != nil {
			select {
			case v, ok := <-ch1:
				if !ok {
					ch1 = nil
					continue
				}
				out <- v
			case v, ok := <-ch2:
				if !ok {
					ch2 = nil
					continue
				}
				out <- v
			}
		}
		close(out)
	}()
	return out
}

func main() {
	ch := fanIn(generator(1), generator(2))
	for v := range ch {
		fmt.Println(v)
	}
}
