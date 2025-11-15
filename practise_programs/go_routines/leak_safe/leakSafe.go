package main

func leak(ch chan int) {
	go func() {
		ch <- 1
		//fmt.Println("sent from leak go routine")
	}()
	//<-ch // leads to blocking of the this go routine as no recever for the go routine channel
}

func safe(ch chan int) {
	go func() {
		ch <- 2
	}()
	<-ch
}

func main() {
	ch := make(chan int)
	leak(ch)
	safe(ch)
}
