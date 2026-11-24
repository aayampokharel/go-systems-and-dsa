package main

import (
	"fmt"
	"sync"
)

func sendInCh(wg *sync.WaitGroup, ch chan<- int, val int) {
	defer wg.Done()
	ch <- val
}

func closeChannel(wg *sync.WaitGroup, ch chan int) {
	wg.Wait()
	close(ch)

}
func main() {
	var ch chan int = make(chan int)
	var wg sync.WaitGroup = sync.WaitGroup{}
	go func() {
		for channelVal := range ch {
			fmt.Println(channelVal)
		}
	}()

	for key := range 100 {
		if key%2 == 0 {
			wg.Add(1)
			go sendInCh(&wg, ch, key)
			wg.Wait()
		}

	}
	go closeChannel(&wg, ch)

	fmt.Println("thankyou")
}
