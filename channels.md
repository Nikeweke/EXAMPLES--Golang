# Channels

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

### GetUser with possible error example

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
