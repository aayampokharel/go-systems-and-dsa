package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID   int
	Data string
}

func Worker(id int, wg *sync.WaitGroup, ch chan Job) {
	defer wg.Done()
	for job := range ch {
		time.Sleep(time.Second * 2)
		fmt.Printf("worker %v processed:::job ID:%v, data:%v\n", id, job.ID, job.Data)
	}

}

func Sender(wg *sync.WaitGroup, ch chan Job, i int) {
	ch <- Job{
		ID:   i,
		Data: fmt.Sprintf("data: %v", i),
	}

}
func main() {
	var wg sync.WaitGroup = sync.WaitGroup{}
	ch := make(chan Job, 10)
	for key := range 3 {
		wg.Add(1)
		go Worker(key, &wg, ch)
	}
	for key := range 5 {
		Sender(&wg, ch, key)
	}
	close(ch)
	wg.Wait()
	fmt.Printf("All workers finished.Working!!")
}
