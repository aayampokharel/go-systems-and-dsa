Write a small main.go file (no external libraries) that does this:

Spawn 3 producer goroutines that send numbers into a single shared channel.

Use a sync.WaitGroup and a dedicated coordinator goroutine to close the channel safely once all 3 producers finish.

Pass a context.WithTimeout (e.g., 50ms) so that if the producers are slow, they exit early via select { case <-ctx.Done(): return } without leaking goroutines.