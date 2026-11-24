package main

import "fmt"

func main() {
	ch := make(chan int)
	for key := range 100 {
		go func() {
			ch <- key
		}()

		select {
		case val, _ := <-ch:
			fmt.Println(val)
			if val == 0 {
				ch = nil
			}
		default:
			fmt.Println("default")
		}

	}
	close(ch)

}
