package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sem := make(chan struct{}, 3)
	wg := sync.WaitGroup{}
	for job := range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			time.Sleep(time.Second * 1)
			fmt.Println(job, "no. completed!")
		}()
	}

	wg.Wait()
	close(sem)
	fmt.Print("thankyou")
}
