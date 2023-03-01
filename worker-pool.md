## Worker pool

* [Article](https://medium.com/code-chasm/go-concurrency-pattern-worker-pool-a437117025b1)
* [More complex solution from article above](https://github.com/syafdia/go-exercise/blob/master/src/concurrency/workerpool/worker_pool.go)

---

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

```go
package main

import (
	"fmt"
	"time"
)

type Response struct {
	Data interface{}
	Error error
	Index int
}

func someOperation(data int) ([]int, error) {
	time.Sleep(time.Second)
	return []int{data}, nil
}



// Convert bi-directional channel to unidirectional channel. - jobs, results
func worker(
	id int, 
	jobs <-chan int, 
	results chan<- Response, 
	fn func(int)([]int, error),
) {
	// loop means that get from channel values
	// like if you do "<-jobs"
	for j := range jobs {
		// fmt.Println("Jobs left:", len(jobs))
		fmt.Println("worker", id, "started  job", j)

		res, err := fn(j)
		fmt.Println("worker", id, "finished job", j)

		results <- Response{ Data: res, Error: err, Index: j }
	}
}

func main() {
	const numJobs = 5
	const workers = 3

	// 1. prepare channels
	jobs := make(chan int, numJobs)
	results := make(chan Response, numJobs)

	// 2. start goroutine, but goroutine wont start till
	// some message appear in channels (worker spawn)
	for w := 1; w <= workers; w++ {
		go worker(
			w, 
			jobs, 
			results,
			someOperation,
		)
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
	close(results)
}

```
