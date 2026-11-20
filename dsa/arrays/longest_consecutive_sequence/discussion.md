```go
for i := 0; i < n; i++ {
    if condition {
        for j := 0; j < n; j++ {
            // operation
        }
    }
}
```


If the condition is true for every iteration, the complexity is \(O(n^2)\).

If it is true for half the iterations, the complexity is still \(O(n^2)\).

If it is true for only one iteration, the complexity is still \(O(n^2)\) because the inner loop executes \(n\) times.

If it is never true, the complexity is \(O(n)\).
