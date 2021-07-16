# Promises

```go
// based on https://madeddu.xyz/posts/go-async-await/
package main

import (
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
)


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

func main() {
	var result1 = <-Promise(func() (interface{}, error) { someComputing(2); return "Promise1: With 2 sec", nil })
	fmt.Println(result1)

	// Promise.all
	var promise1 = Promise(func() (interface{}, error) { return getData() })
	var promise2 = Promise(func() (interface{}, error) { return getData() })
	result1, result2 := <-promise1, <-promise2
	fmt.Println(string(result1.Result.([]byte)), string(result2.Result.([]byte)))


	// Promise.race
	var result3 PromiseResponse
	var promise3 = Promise(func() (interface{}, error) { someComputing(2); return "PromiseRace: With 2 sec", nil })
	var promise4 = Promise(func() (interface{}, error) { someComputing(4); return "PromiseRace: With 4 sec", nil })
	select {
		case result3 = <-promise3:
		case result3 = <-promise4:
	}
	fmt.Println(result3)
}
```
