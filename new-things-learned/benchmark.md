- for testing using benchmark i need to have prefix of the function starting with `Benchmark` exactly. 
- with modern `b.Loop()`, `resetTimer()` method is redundant,Go's modern `b.Loop()` syntax (introduced in Go 1.24), the framework automatically isolates and measures only the code inside the loop.


### How Benchmark Works


```text
BenchmarkContainsDuplicate/size_of_for_generateSlice:10-20         11413054            95.01 ns/op            0 B/op        0 allocs/op
...
PASS 

```

Two major red flags stood out:

1. **`0 B/op` and `0 allocs/op`:** Both `hasDuplicate` and `BestSolution` rely on Go `map` objects. Maps inherently allocate memory on the heap. Seeing zero allocations meant the map wasn't actually running or the compiler optimized the code away entirely.
2. **The Output Label (`generateSlice`):** The benchmark label was measuring `generateSlice` inside the timed loop instead of measuring `hasDuplicate` directly.


`allocs/op` counts the number of individual times memory is requested from the heap, not the size of the memory itself.

The distinction between total bytes `(B/op)` and allocation counts `(allocs/op)` works as follows:

- *One Large Allocation*: Allocating a single 25 MB slice all at once results in 1 alloc/op and 25,000,000 B/op.

- *Many Small Allocations*: Allocating 25 million individual 1-byte variables results in 25,000,000 allocs/op and 25,000,000 B/op.

Even though the total byte size (B/op) is identical in both scenarios, the second scenario destroys performance. Each allocation call (alloc) introduces CPU overhead—searching for free space in the heap, managing pointers, and adding work for the Garbage Collector.

---

### Breaking Down the Benchmark Metrics (`ns/op`, `B/op`, etc.)

When running `go test -bench=. -benchmem`, Go outputs performance data broken down like this:

* **`11413054`** (Iterations): How many times Go was able to run the benchmark loop during the standard test period (usually 1 second). Higher numbers mean faster code.
* **`95.01 ns/op`** (Nanoseconds per Operation): The average time it took to complete **one single execution** of the tested function, measured in nanoseconds ($1 \text{ millisecond} = 1,000,000 \text{ nanoseconds}$). Lower is faster.
* **`0 B/op`** (Bytes per Operation): The average amount of heap memory allocated **per run** of the function, measured in bytes.
* **`0 allocs/op`** (Allocations per Operation): How many distinct times memory was requested from the heap **per run**. Keeping this as low as possible reduces pressure on Go's Garbage Collector.