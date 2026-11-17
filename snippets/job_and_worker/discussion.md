It is completely unpredictable (non-deterministic) which worker will pick up which job.

In Go, `for job := range ch` does not mean "loop through whatever is currently in the channel and then stop." Instead, it means:

- Read items as long as they are available.
- If the channel is currently empty, block (go to sleep) and wait for new items to arrive.
- Keep doing this indefinitely... unless the channel is closed.

---
### What happens if I forget to close(ch)?
If I omitted close(ch) at the end of main, my program would experience a deadlock or a goroutine leak:

The workers would finish processing the 5 jobs.

Once the channel buffer became empty, the workers would go back to the top of the loop (for job := range ch) and block forever, waiting for jobs that are never going to come.

Because the workers never exit, wg.Done() is never called, and wg.Wait() in main() hangs forever.


---