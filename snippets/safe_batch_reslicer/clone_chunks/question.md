Part 1: CloneChunks(items []int, chunkSize int) [][]int
Split items into contiguous chunks of up to chunkSize elements (for example, items = []int{14, -7, 38, 2, 91, -25, 19} with chunkSize = 3 produces chunks [14, -7, 38], [2, 91, -25], and [19]).

Edge cases: If chunkSize <= 0 or len(items) == 0, return nil.

Backing-array isolation: Every returned chunk must own its own underlying array (make + copy) so mutating a returned chunk after CloneChunks returns never mutates items or any sibling chunk.

Part 2: 3-Worker Pool & Overall Maximum in main()
Run a producer goroutine that loops over the chunks and sends each []int chunk one by one onto a jobs channel, then closes jobs once all chunks are sent.

Launch 3 worker goroutines (numWorkers := 3) that read chunks using for chunk := range jobs, find the largest integer in that chunk, and send each chunk's maximum to a results channel.

Use a sync.WaitGroup to wait until all 3 workers finish, then close results.

In main(), range over results to compute the overall maximum across all chunks and print it (and print a clear message such as "no valid chunks" if the result is empty/nil).