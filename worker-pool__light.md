## Worker pool with no limit on workers

```go
package main

import (
	"fmt"
	"time"
)

type Fn func() (interface{}, error)
type Response struct {
	Data interface{}
	Error error
}

// its function wrapper that set passed-in function into goroutine using channel for get response
func New(results chan<- Response, f Fn) { 
	go func() { 
		result, err := f() 
		results <- Response{ Data: result, Error: err }
	}() 
}

func main() {
	var resultsChan = make(chan Response)
	var results = []Response{}
	var jobs = []int{2, 1 ,4} // some data that need to be passed to function
	var jobsLen = len(jobs)

	for _, i := range jobs {
		i := i

		go func() { 
			result, err := someComputing(i) 
			resultsChan <- Response{ Data: result, Error: err }
		}() 

		// New(resultsChan, func() (interface{}, error) { 
		// 	return someComputing(i); 
		// })
	}

	for j := 1; j <= jobsLen; j++ {
		var res = <-resultsChan
		fmt.Println(res)
		results = append(results, res)
	}
	close(resultsChan)

  fmt.Println(results)
}

func someComputing(delay int) (string, error) {
	time.Sleep(time.Second*time.Duration(delay))
	return fmt.Sprintf("result %d", delay), nil
}
```
