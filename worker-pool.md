## Worker pool

```go
package main

import (
	"fmt"
	"time"
)

// Convert bi-directional channel to unidirectional channel. - jobs, results
func worker(id int, jobs <-chan int, results chan<- int) {
	
	// loop means that get from channel values
	// like if you do "<-jobs"
	for j := range jobs {
		fmt.Println("Jobs left:", len(jobs))
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)

		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	const workers = 3

	// 1. prepare channels
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 2. start goroutine, but goroutine wont start till
	// some message appear in channels (worker spawn)
	for w := 1; w <= workers; w++ {
		go worker(w, jobs, results)
	}

	// 3. adding messages into jobs channel, 
	// it will trigger goroutines above (add jobs in channel)
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 4. read results 
	for a := 1; a <= numJobs; a++ {
		// <-results
		fmt.Println("Result", <-results)
	}
}

```
