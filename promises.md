# Promises 

> BEtter to use worker-pool https://github.com/Nikeweke/EXAMPLES--Golang/blob/master/worker-pool__light.md

> place this file into separate file, for example: /services/promise/index.go to keep package name

```go
package promise

type Fn func() (interface{}, error)
type Pending <-chan Response 
type Response struct {
	Data interface{}
	Error error
}

// its function wrapper that set passed-in function into goroutine using channel for get response
func New(f Fn) Pending { 
	var result interface{} 
	var err error 
	
	c := make(chan Response)
	go func() { 
		defer close(c) 
		result, err = f() 
		c <- Response{ Data: result, Error: err }
	}() 
	
	return c 
}

// just loop over channels("[]PromisePending") and wait them to end, 
// and get responses("[]PromiseResponse") from all goroutines
func All(promises []Pending) []Response {
	var results = []Response{}
	for _, promise := range promises {
		results = append(results, <-promise)
	}
	return results
}

```
<br />

###### Await one 
```go
  var result1 = <-promise.New(func() (interface{}, error) { 
		return someComputing(2); 
	})
  fmt.Println(result1)
```
<br />

###### Await a few
```go
// Await a few
var promise1 = promise.New(func() (interface{}, error) { 
	return someComputing(2); 
})
var promise2 = promise.New(func() (interface{}, error) { 
	return someComputing(4); 
})
result1, result2 := <-promise1, <-promise2
fmt.Println(result1.Data, result2.Data)
```
<br />


###### Promise.all with array 
```go
var promises := []promise.Pending{
	promise.New(func() (interface{}, error) {
		user, err := models.User{} 
			err := DB.
				Where("id = 1").
				First(&user).
				Error
			return user, err
	}),
	
	promise.New(func() (interface{}, error) {
		user, err := models.User{} 
			err := DB.
				Where("id = 2").
				First(&user).
				Error
			return user, err
	}),
}
var results = promise.All(promises) // returns []PromiseResponse{ Data, Error }
fmt.Println(results[0].Data) // or you can cast results[0].Data.(string) or .([some-type])
```
<br />


###### Promise.all with array witn passing in arguments + loop
```go
var promises = []promise.Pending{}
var data = []string{"123", "234"}

// fnWrapper - here is for passing "item" param into function (you can go without this, just call as anonym function),
// without this wrapper function will always output only last item from array "data" (closure stuff)
var fnWrapper = func(item string) promise.Fn {
	return func() (interface{}, error) { 
		fmt.Println(item)
		return someComputing(2); 
	}
}

for _, item := range data {
	promises = append(
		promises, 
		promise.New(fnWrapper(item)),
	)
}

results := promise.All(promises)
fmt.Println(results)
```
<br />

###### Promise.race
```go
// Promise.race
var result3 promise.Response
var promise3 = Promise(func() (interface{}, error) { return someComputing(2) })
var promise4 = Promise(func() (interface{}, error) { return someComputing(2) })
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
func someComputing(delay time.Duration) (string, error) {
	time.Sleep(time.Second*delay)
	return "Some result", nil
}
```
