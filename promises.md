# Promises

```go
type PromiseFn func() (interface{}, error)
type PromisePending <-chan PromiseResponse 
type PromiseResponse struct {
	Data interface{}
	Error error
}

// its function wrapper that set passed-in function into goroutine using channel for get response
func NewPromise(f PromiseFn) PromisePending { 
	var result interface{} 
	var err error 
	
	c := make(chan PromiseResponse)
	go func() { 
		defer close(c) 
		result, err = f() 
		c <- PromiseResponse{ Data: result, Error: err }
	}() 
	
	return c 
}

// just loop over channels("[]PromisePending") and wait them to end, 
// and get responses("[]PromiseResponse") from all goroutines
func PromiseAll(promises []PromisePending) []PromiseResponse {
	var results = []PromiseResponse{}
	for _, promise := range promises {
		results = append(results, <-promise)
	}
	return results
}

```
<br />

###### Await one 
```go
  var result1 = <-NewPromise(func() (interface{}, error) { someComputing(2); return "Promise1: With 2 sec", nil })
  fmt.Println(result1)
```
<br />

###### Await a few
```go
// Await a few
var promise1 = NewPromise(func() (interface{}, error) { 
	someComputing(2); 
	return "PromiseAll: With 2 sec", nil 
})
var promise2 = NewPromise(func() (interface{}, error) { 
	someComputing(4); 
	return "PromiseAll: With 4 sec", nil  
})
result1, result2 := <-promise1, <-promise2
fmt.Println(result1.Data, result2.Data)
```
<br />


###### Promise.all with array 
```go
var promises := []PromisePending{
	NewPromise(func() (interface{}, error) {
		user, err := models.User{} 
			err := DB.
				Where("id = 1").
				First(&user).
				Error
			return user, err
	}),
	
	NewPromise(func() (interface{}, error) {
		user, err := models.User{} 
			err := DB.
				Where("id = 2").
				First(&user).
				Error
			return user, err
	}),
}
var results = PromiseAll(promises) // returns []PromiseResponse{ Data, Error }
fmt.Println(results[0].Data)
```
<br />


###### Promise.all with array witn passing in arguments + loop
```go
var promises = []PromisePending{}
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
	}
	var promise = Promise(fnWrapper(item))
	promises = append(promises, promise)
}

results := PromiseAll(promises)
fmt.Println(results)
```
<br />

###### Promise.race
```go
// Promise.race
var result3 PromiseResponse
var promise3 = Promise(func() (interface{}, error) { someComputing(2); return "PromiseRace: With 2 sec", nil })
var promise4 = Promise(func() (interface{}, error) { someComputing(4); return "PromiseRace: With 4 sec", nil })
select {
	case result3 = <-promise3:
	case result3 = <-promise4:
}
fmt.Println(result3)
```
<br />


--- 


### someComputing fn
```go
func someComputing(delay time.Duration) {
	time.Sleep(time.Second*delay)
}
```
