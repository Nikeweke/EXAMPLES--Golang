# Channels

> https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb

### Basic
```go
package main

import (
  "fmt"
)

func main() {
  // create a channel 
  channel := make(chan User) // if "make(chane int, 100)" its bufferized channel 

  // WRITIING into channel 
  // (without "go func" it will be error - " all goroutines are asleep - deadlock!")
  go func() {
    channel <- User{"Alex Bradvik", 21}
    channel <- User{"Jessica Avaro", 40}
  }()

  // READING from channel (awaiting data)
  for i := 0; i < 2; i++ {
    data:= <-channel
    fmt.Println(data)
  }

}
```
<br />

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
