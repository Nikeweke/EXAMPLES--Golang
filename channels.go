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
