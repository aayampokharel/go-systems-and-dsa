package main

import (
	"fmt"
	"sync"
)

func Producer(wg *sync.WaitGroup, ch chan<- int, val int) {
	defer wg.Done()
	ch <- val
}

func Consumer(ch <-chan int) {
	fmt.Println(<-ch)
	fmt.Println("consumer executed")
}

func closeChannel(wg *sync.WaitGroup, ch chan int) {
	wg.Wait()
	close(ch)
	fmt.Println("channel closed")

}

func main() {
	var wg = sync.WaitGroup{}
	var ch chan int = make(chan int)
	for key := range 5 {
		wg.Add(1)
		go closeChannel(&wg, ch)
		go Producer(&wg, ch, key)
		go Consumer(ch)
	}
	wg.Wait()
	fmt.Print("goroutine example completed")
}
