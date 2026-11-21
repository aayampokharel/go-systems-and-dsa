In Go, Producer–Consumer and Worker Pool are concurrency patterns commonly implemented using goroutines and channels.


First, understand that these are not three competing concepts. They describe different responsibilities in a concurrent system.
- Producer: Generates work (data or tasks).
- Consumer: Receives and processes work.
- Worker: Performs a particular unit of work. A worker is usually a goroutine.


1. Producer pattern
A producer is responsible for generating data and sending it somewhere.

Understanding these two things:
```go
func producer(ch chan<- int)
```

chan<- int means this function can only send integers into the channel.

2. Consumer pattern
A consumer receives data and processes it.

Notice the difference:
```go
func producer(ch chan<- int) //Can only send.
func consumer(ch <-chan int) //can only receive
```


3. Worker pattern
 A worker is simply a goroutine that performs work.


# job is simply a unit of work that needs to be performed.
The channel is just the transportation mechanism. The jobs are the actual data being transported.

---

