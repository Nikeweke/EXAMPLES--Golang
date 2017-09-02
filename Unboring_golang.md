## Golang 

#### Содержание


* [Variables](#variables)
* [Arrays](#arrays)
* [Functions](#functions)
* [For](#for)
* [While](#while)
* [If, switch](#if-switch)
* [Defer](#defer)
---

### Variables

###### ВЫВОД типа переменной
```go
fmt.Printf("i is of type %T\n", i) // out-> i is of type int
```

```go
var x int
var s string
var d int32 
var b float32
var q float64

x    := "Hello"
var x = "Hello"

x    := 5
var x = 5

var c, python, java = true, false, "no!"
var i, j int        = 1, 2

//  ВЫВОД типа переменной
fmt.Printf("i is of type %T\n", i) // out-> i is of type int
```


### Arrays

```go
primes := [6]int{2, 3, 5, 7, 11, 13}
words  := [3]string{"Hello", "mate", "world"}
```


### Functions
###### Возвращение нескольких результатов
```go
package main
import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

```

###### Возвращение названых переменных
```go
package main
import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```

### For
##### FOR обычный
```go

 sum := 0
for i := 0; i < 10; i++ {
 sum += i
}
fmt.Println(sum)
```

##### FOR без итератора
```go
sum := 1
for ; sum < 1000; {
  sum += sum
}
fmt.Println(sum)
```

##### FOR без конца и края
```go
for{
  fmt.Println("Forever hello world")
}
```


### While
```go
sum := 1
for sum < 1000 {
  sum += sum
}
fmt.Println(sum)
```

### If, switch
##### If 
```go
func pow(x, n, lim float64) float64 {

	if v := math.Pow(x, n); v < lim { // вот тут сначала исполняеться функция math.Pow(),
                                      // и результат идет в "v", а потот юзаеться в сравненнии "v < lim" 
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

##### Switch обычный 
```go
fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
```

##### Switch без условия. 
Если условия нет, значит идет как "true". Это можно использовать вместо длинных if - elseif - else
```go
t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
```

### Defer
defer - выполниться после достижение и окончания последней строки ф-ции в которой она написана
```go
defer fmt.Println("world")
fmt.Println("hello")
```

Также можно накидывать в очередь выполнения defer (отработает по типу last-in-first-out)
```go
fmt.Println("counting")
for i := 0; i < 10; i++ {
 defer fmt.Println(i)
}
fmt.Println("done")
```
