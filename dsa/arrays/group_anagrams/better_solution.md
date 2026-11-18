# Why `BetterApproach` is better

- My first solution sorts every string: converting to `[]rune`, sorting, then building a string key. That costs **O(m log m)** per string plus extra allocations.
- `BetterApproach` counts characters into a fixed `[26]byte` and uses that array as the map key. Counting is **O(m)** and needs no sorting or extra string allocation.
- Arrays are comparable in Go, so `[26]byte` works as a map key just like a string.
- It also preallocates the result slice with `make([][]string, 0, len(m))`.

Overall: **O(n·m)** instead of **O(n·m log m)**. The trade-off is that it only works for lowercase `a-z`, which is exactly what the problem guarantees.
