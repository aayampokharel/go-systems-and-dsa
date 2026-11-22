```go
s:="123456"
s=append(s[:mid],s[mid+1:]...)
```

if s is string, append doesnot work since first param for append should be slice and slice is immutable.
#### Convert it to a byte slice

