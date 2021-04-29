# Channels and goroutines

> * https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb
> * https://madeddu.xyz/posts/go-async-await/

* [Basic](#Basic)
* [Buffered size channel](#Buffered-size-channel)
* [Channel length and capacity](#Channel-length-and-capacity)
* [Unidirectional channels](#Unidirectional-channels)
* [Select](#Select)
* ["GetUser" with error in return - one iteration](#getuser-with-error-in-return---one-iteration)
* ["GetUser" with error in return  - few iteration](#getuser-with-error-in-return----few-iteration)

--- 

### Basic
```go
package main
import "fmt"

func main() {
  // create a channel 
  channel := make(chan string)

  // WRITIING into channel 
  // (without "go func" it will be error - " all goroutines are asleep - deadlock!")
  go func() {
    channel <- "Alex Bradvik"
    channel <- "Jessica Avaro"
  }()

  // READING from channel (awaiting data)
  for i := 0; i < 2; i++ {
    data:= <-channel
    fmt.Println(data)
  }

  // ----> OR
  // result1, result2 := <-channel, <-channel
  // fmt.Println(result1)
  // fmt.Println(result2) 
}
```
<br />

### Buffered size channel
If we will write a few values into unbuffered channel it will be deadlock

```go
func greet(ch chan string) {
  for i:=0; i < 3; i++ {
    fmt.Println("Hello "+ <-ch)
  }
}

func main() {
  var channel = make(chan string)

  go greet(channel)

  channel <- "Jake" // "Hello Jake"
  channel <- "Alex" // ERROR: deadlock

  fmt.Println("done")
}
```

With buffered size we can do it. When channel overflowed - goroutine will launch
```go 
func greet(ch chan string) {
  for i:=0; i < 3; i++ {
    fmt.Println("Hello "+ <-ch)
  }
}

func main() {
  // 1. set size of channel
  var channel = make(chan string, 3)

  // 2. start goroutine
  go greet(channel)

  // 3. push into channel values
  channel <- "Jake"
  channel <- "Alex"
  channel <- "Alex2"

  // 4. on purpose overflowing channel will cause to goroutine execution
  channel <- "" 

  fmt.Println("done")
}
```

Also we can get from closed channel items cuz they are buffered
```go
func main() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	
	// iteration terminates after receving 3 values
	for elem := range c {
		fmt.Println(elem)
	}
}
```
<br />

### Channel length and capacity
```go
func main() {
  c := make(chan int, 3)
	c <- 1
	c <- 2

  // capacity and length will work only with buffered size channel
  // cuz unbuff channel need to cover with goroutine above

	fmt.Printf("Length of channel c is %v and capacity of channel c is %v", len(c), cap(c))
	fmt.Println()
}
```
<br />

### Unidirectional channels

Using unidirectional channels increases the type-safety of a program. Hence the program is less prone to error.

```go
package main
import "fmt"

func main() {
	roc := make(<-chan int) // receive-only channel
	soc := make(chan<- int) // send-only channel

	fmt.Printf("Data type of roc is `%T`\n", roc)
	fmt.Printf("Data type of soc is `%T\n", soc)
}
```

Convert bi-directional channel to unidirectional channel.
```go

package main
import "fmt"

func greet(roc <-chan string) {
	fmt.Println("Hello " + <-roc + "!")
}

func main() {
	fmt.Println("main() started")
	c := make(chan string)

	go greet(c)

	c <- "John"
	fmt.Println("main() stopped")
} 
```
<br />


### Select
Works like `switch` but for channels

```go
package main

import (
	"fmt"
	"time"
)

var start time.Time
func init() {
	start = time.Now()
}

func service1(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello from service 1"
}

func service2(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "Hello from service 2"
}

func main() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	}
	
	fmt.Println("main() stopped", time.Since(start))
}
```


### "GetUser" with error in return - one iteration

```go
package main

import (
  "fmt"
  "errors"
  "time"
)

var (
	ErrPersonNotFound = errors.New("account_not_found")
)

type User struct {
  Name string
  Age int
}

var Users = map[int]User{
  21: User{"Alex Bradvik", 21},
  40: User{"Jessica Avaro", 40},
}

func main() {
  // create a channel 
  var channel = make(chan User) 
  var errCh = make(chan error)
  var userId = 22

  // getting error that can function return into error-channel
  go func() {  errCh <- GetUser(userId, channel) }()

  // listening for response from corresponding channel
  select {
    case msg1 := <-channel:  
    fmt.Println("Success: ", msg1)

    case msg2 := <-errCh:  
    fmt.Println("Error found:", msg2)
  }
}

func GetUser(id int, out chan User) error {
	var result = Users[id]
	if result == (User{}) {
    return ErrPersonNotFound
	}
	time.Sleep(time.Millisecond * 2000)
	out <- result
  return nil
}
```
<br />

### "GetUser" with error in return  - few iteration

```go
package main

import (
  "fmt"
  "errors"
  "time"
  // "sync"
)

var (
	ErrPersonNotFound = errors.New("account_not_found")
)

type User struct {
  Name string
  Age int
}

var Users = map[int]User{
  21: User{"Alex Bradvik", 21},
  40: User{"Jessica Avaro", 40},
}

func main() {
  // create a channels 
  var channel = make(chan User) 
  var errCh = make(chan error)

  var userId = 21
  var iterations = 10

  // make parallel a few iterations to get user 
  for i:=0; i < iterations; i++ {
    go func(userId int) {
      errCh <- GetUser(userId, channel) 
    }(userId)
  }

  var result = []User{}

  // wait for result 
  for i := 0; i < iterations; i++ {

    // "select" listening for values from channel 
    select {
      // i set here check cuz sometime it can trapped into error case for some reason
      case err := <-errCh:  
        if err != nil {
          fmt.Println("Error found:", err)
          continue
        }
        result = append(result, <-channel)

      case user := <-channel:  
        fmt.Println("Success: ", user)
        result = append(result, user)
    }
  }


  fmt.Println(len(result))
}

func GetUser(id int, out chan User) error {
  fmt.Println("HERE")
	var result = Users[id]
	if result == (User{}) {
    return ErrPersonNotFound
	}
	time.Sleep(time.Millisecond * 2000)
	out <- result
  return nil
}
```


