# Why `BestSolution` is better

- My first solution scans the whole frequency map **k times** to find the current maximum, then deletes it. That is **O(k·d)** time (where `d` is the number of distinct values), which becomes **O(n²)** when `k` is close to `n`.
- `BestSolution` counts frequencies once, then puts each value into a bucket indexed by its frequency: `buckets[freq]`. Frequency is bounded by `len(nums)`, so the bucket array has size `n+1`.
- Walking the buckets from highest frequency to lowest and collecting until `k` elements gives **O(n + d + n) = O(n)** time and **O(n)** space.
- It avoids sorting the frequencies (**O(d log d)**) and avoids mutating/deleting from the map while searching.

Overall: **O(n)** instead of **O(k·d)** / **O(n²)**.
