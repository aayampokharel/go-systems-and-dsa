The Requirements:
Write a small Go program (main.go) that implements a worker pool:

A Job struct containing an ID (integer) and some data (e.g., string).

A worker function that runs concurrently in a goroutine, reads jobs from a shared inbound channel (chan Job), processes them (simulate with a brief time.Sleep), and prints out which worker processed which job.

In main():

Create a buffered jobs channel (e.g., capacity 10).

Spawn 3 worker goroutines using a sync.WaitGroup.

Send 5 jobs into the jobs channel from the main goroutine.

Close the jobs channel once all jobs are sent so the workers know to stop blocking and exit cleanly.

Wait for all workers to finish using wg.Wait(), then print "All workers finished. Exiting."