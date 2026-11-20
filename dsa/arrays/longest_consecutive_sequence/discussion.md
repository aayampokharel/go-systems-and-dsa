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


## How to identify this in future problems?
Remember three questions while encountering nested loops:
1. Can the inner loop execute \(n\) times for every outer iteration?
2. Can the same elements be processed repeatedly by different outer iterations?
3. Is there some condition that prevents repeated processing?
If the same elements are repeatedly processed, quadratic complexity might occur.

If each element can be processed only a constant number of times across the entire algorithm, this might have \(O(n)\) complexity.
This is closely related to amortized analysis.