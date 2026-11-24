nilCh <- 1 
triggers first and causes a deadlock (a fatal runtime error), not a panic, so recover() will not catch it in main.Code is below: 


```go
func main() {
	var nilCh chan int
	var closedCh chan int = make(chan int)
	close(closedCh)
	var workingCh chan int = make(chan int)

	nilCh <- 1
	closedCh <- 1
	workingCh <- 1
}
```

## output: 
![alt text](image.png)


--
# IF I SEND TO CLOSED CHANNEL: 

```go
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic:%v", r)
		}
	}()
	var closedCh chan int = make(chan int)
	close(closedCh)
	closedCh <- 1
}
```
##  OUTPUT: 

```bash
╰─$ go run . 
panic:send on closed channel% 
```

---

```go

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
```

>[!important] the solution here would be to: fan-in pattern: using waitgroups here to know if all sending channels have sent their stuffs or not, and then close the channel; while in receiving side: I simply use the for...range loop over it. 


---
