package main

import (
	"fmt"
	"math"
	"slices"
	"sync"
)

func closeChannel(wg *sync.WaitGroup, jobs chan []int) {
	wg.Wait()
	close(jobs)
}

func receiveResults(receiveResultChan chan int, done chan struct{}) {
	if receiveResultChan == nil {
		fmt.Println("nil channel")
		return
	}
	highVal := math.MinInt
	for valInt := range receiveResultChan {

		if valInt > highVal {
			highVal = valInt
		}

	}
	fmt.Println("max value is:: ", highVal)
	done <- struct{}{}
}

func CloneChunks(items []int, chunkSize int) [][]int {
	var return2DSlice [][]int
	if chunkSize <= 0 {
		chunkSize = 1
	}
	var tempItemsOfChunkSizeMax []int

	for i := 0; i < len(items); i += chunkSize {
		tempItemsOfChunkSizeMax = []int{}
		if i+chunkSize <= len(items) {
			tempItemsOfChunkSizeMax = append(tempItemsOfChunkSizeMax, items[i:i+chunkSize]...)

		} else {
			tempItemsOfChunkSizeMax = append(tempItemsOfChunkSizeMax, items[i:]...)
		}
		return2DSlice = append(return2DSlice, tempItemsOfChunkSizeMax)
	}
	return return2DSlice
}

func Producer(wg *sync.WaitGroup, jobs chan<- []int, value []int, continueSpawnning chan struct{}) {
	defer wg.Done()
	defer func() {
		<-continueSpawnning
	}()
	jobs <- value

}

func main() {
	workersNumber := 3

	continueSpawnning := make(chan struct{}, workersNumber)
	var wg *sync.WaitGroup = &sync.WaitGroup{}
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 29, 31, 32, 33, 34, 345}
	chunkSize := 1
	twoDSlice := CloneChunks(items, chunkSize)
	jobs, result, done := make(chan []int, len(twoDSlice)), make(chan int), make(chan struct{})

	for _, sliceAsVal := range twoDSlice {
		continueSpawnning <- struct{}{}
		wg.Add(1)
		go Producer(wg, jobs, sliceAsVal, continueSpawnning)
	}
	go closeChannel(wg, jobs)
	go receiveResults(result, done)

	for sliceAsVal := range jobs {
		slices.Sort(sliceAsVal)
		result <- sliceAsVal[len(sliceAsVal)-1]

	}
	close(result)
	<-done

}
