## Проблемы и их решение

1. **FOR LOOP**. Есть массив типа - `map[User]string` , где `User` структура в которой одне поле **Id string**. Надо достать значение которое справу (value). При обычном ренже выдает значение только ключа(key). Решение:
```golang

type User struct {
	Id       string  
}

func main(){ 

var clients = make(map[User]string)

...

// так выведет значение ключа(key)
for client := range clients {
   fmt.Println(client)
   fmt.Println(client.Id)
}

// так выведет значение справа(value)
for _, client := range clients {
   fmt.Println(client)
}

// так выведет оба значения(key & value)
for i, client := range clients {
   fmt.Println(i.Id)
   fmt.Println(client)
}
```

2. **CHANNELS**. Каналы работают только в goroutines (потоках) иначае ошибка - deadlock. 
```golang
package main
import "fmt"

//MAIN
func main() {
    c := make(chan int, 100)
    
    // Вот так не будет работать
    c <- 1
    c <- 2
    numb := <-c
    fmt.Prinln(numb) // deadlock !
    
    // Вот так работает
    go func(){
      c <- 2
      c <- 1
      c <- 4
      c <- 5
   }()

    fmt.Println(<-c)
    fmt.Println("THAT LEN",  len(c))

    fmt.Println(<-c)
    fmt.Println("THAT LEN",len(c))
    fmt.Println(<-c)
    fmt.Println(<-c)
    
    fmt.Scanln()
}

```
