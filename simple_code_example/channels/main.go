package main

import (
	"fmt"
	"runtime"
	"time"
)

func fastAPI(ch chan<- int) {
	time.Sleep(time.Millisecond * 1)
	ch <- 1
}
func slowAPI(ch chan<- int) {
	time.Sleep(time.Millisecond * 100)
	ch <- 100
}
func main() {
	var ch chan int = make(chan int)
	for range 100 {

		go fastAPI(ch)
		go slowAPI(ch)
		fmt.Println(<-ch, "no. of go routines", runtime.NumGoroutine())
	}

}
