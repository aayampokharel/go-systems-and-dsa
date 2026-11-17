package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func sender(ctx context.Context, wg *sync.WaitGroup, ch chan int, i int) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		fmt.Println("ctx timeout for sender:::", ctx.Err())
		return
	case ch <- i * 10:
		fmt.Println("value sent")
	}
}

func closeChan(wg *sync.WaitGroup, ch chan int) {
	wg.Wait()
	close(ch)
	fmt.Println("All producers finished. Channel closed.")
}

func main() {
	var wg *sync.WaitGroup = &sync.WaitGroup{}
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	for key := range 3 {
		wg.Add(1)
		go sender(ctx, wg, ch, key)
		time.Sleep(time.Second * 1)
	}
	go closeChan(wg, ch)
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Println("channels already closed")
				return
			} else {
				fmt.Println(val)
			}
		case <-ctx.Done():
			fmt.Println("timeout:", ctx.Err())
			return
		}
	}

}
