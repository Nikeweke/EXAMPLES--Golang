# Tips

### Запись в канал без запуска goroutine приведет к ошибке

```go
func greet(ch chan string) {
  fmt.Println("Hello "+ <-ch)
}

func main() {
  var channel = make(chan string)

  // go greet(channel)

  // push into channel without launching previosly goroutine(go greet()) will lead to CRASH!
  channel <- "Jake"

  fmt.Println("done")
}
```
<br />

### После запуска goroutine в канало надо отправить что-то в канал чтобы её разбудить

```go
func greet(ch chan string) {
  // без задержки не успеет прочитать из канала, 
  // но если отправку в канал поставить сразу после запуска горутины, то можно убрать задержку
  time.Sleep(time.Millisecond*5) 
  fmt.Println("Hello "+ <-ch)
}

func main() {
  var channel = make(chan string)

  go greet(channel)

  fmt.Println("done")
  channel <- "Jake"

  fmt.Println("done")
}
```
<br />

### Запуск горутины и несколько раз запись в канал
```go

func greet(ch chan string) {
  fmt.Println("Hello "+ <-ch)
}

func main() {
  var channel = make(chan string)

  go greet(channel)

  channel <- "Jake" // Hello jake
  channel <- "Alex" // Error: Deadlock

  fmt.Println("done")
}
```
<br />
