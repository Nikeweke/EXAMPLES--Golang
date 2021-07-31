# Promises, async/await

```go
// based on https://madeddu.xyz/posts/go-async-await/
package main

import (
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
)

func main() {
	// Await one 
	var result1 = <-Promise(func() (interface{}, error) { someComputing(2); return "Promise1: With 2 sec", nil })
	fmt.Println(result1)


	// Promise.all
	var promise1 = Promise(func() (interface{}, error) { someComputing(2); return "PromiseAll: With 2 sec", nil })
	var promise2 = Promise(func() (interface{}, error) { someComputing(4); return "PromiseAll: With 4 sec", nil  })
	result1, result2 := <-promise1, <-promise2
	fmt.Println(result1.Result, result2.Result)


	// Promise.race
	var result3 PromiseResponse
	var promise3 = Promise(func() (interface{}, error) { someComputing(2); return "PromiseRace: With 2 sec", nil })
	var promise4 = Promise(func() (interface{}, error) { someComputing(4); return "PromiseRace: With 4 sec", nil })
	select {
		case result3 = <-promise3:
		case result3 = <-promise4:
	}
	fmt.Println(result3)

	
	// Promises Array
	var results = []PromiseResponse{}
	var promises = []<-chan PromiseResponse{}
	var data = []string{"123", "234"}

	for _, item := range data {
		// fnWrapper - here is for passing "item" param into function (you can go without this, just call as anonym function),
		// without this wrapper function will always output only last item from array "data"
		var fnWrapper = func(item string) func() (interface{}, error) {
			return func() (interface{}, error) { 
				fmt.Println(item)
				someComputing(2); 
				return "Promise1: With 2 sec", nil 
			}
		}(item)
		var promise = Promise(fnWrapper(item))
		promises = append(promises, promise)
	}
	

	for _, promise := range promises {
		results = append(results, <-promise)
	}
	fmt.Println(results)
}


// ========================================> PROMISE Itself
type PromiseResponse struct {
	Result interface{}
	Error error
}

func Promise(f func() (interface{}, error)) <-chan PromiseResponse { 
	var result interface{} 
	var err error 
	
	c := make(chan PromiseResponse)
	go func() { 
		defer close(c) 
		result, err = f() 
		c <- PromiseResponse{ Result: result, Error: err }
	}() 
	
	return c 
}


// ========================================> EXAMPLE FNs
func getData() ([]byte, error) {
	url := "https://apip.parts-point.com/" 
	resp, err := http.Get(url) 
	if err != nil { 
		return nil, err 
	} 
		
	defer resp.Body.Close() 
		
	return ioutil.ReadAll(resp.Body) 
}

func someComputing(delay time.Duration) {
	time.Sleep(time.Second*delay)
}


```
