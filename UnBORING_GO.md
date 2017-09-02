## Golang 

#### Содержание


* [Variables](#variables)
* [Arrays, Slices](#arrays)
* [Maps](#maps)
* [Functions](#functions)
* [For](#for)
* [While](#while)
* [Range](#range)
* [If, switch](#if-switch)
* [Defer](#defer)
* [Structs (alike Class or Object)](#)
---

### Variables

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
```

###### ВЫВОД типа переменной
```go
fmt.Printf("i is of type %T\n", i) // out-> i is of type int
```


### Arrays, Slices
###### Array(Массив) - имеет установленный размер
```go
primes := [6]int{2, 3, 5, 7, 11, 13}
words  := [3]string{"Hello", "mate", "world"}
```

###### Slices(Срезы) - массив с динамическим размером и более гибкий
```go
q := []int{2, 3, 5, 7, 11, 13}
r := []bool{true, false, true, true, false, true}
a := make([]int, 5)  // len(a)=5

fmt.Println(q[0:3]) // out-> 2, 3, 5

len(q) // длина
cap(q) // вместимое

```

###### Slice пустой == nil
```go
var s []int
fmt.Println(s, len(s), cap(s))
if s == nil {
	fmt.Println("nil!")
}
```

###### Slice добавить в срез
```go
var s []int

// append works on nil slices.
s = append(s, 0)
s = append(s, 2, 3, 4)
fmt.Println(s) // out-> 0, 2, 3, 4
```

### Maps

###### Создание карты
```go
// Создание карты
elements := map[string]string{
  "H": "Hydrogen",
}
// OR
elements := make(map[sting]string)
```

###### Создание нового элемента
```go
elements["Li"] = "Lithium"   // +1 элемент в - elements
elements["Be"] = "Berrylium" // +1 элемент в - elements
fmt.Println(elements)
```

###### Проверка на существование эл. в карте
```go
name, ok := elements["He"]
fmt.Println(name, ok)
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
###### FOR обычный
```go
sum := 0
for i := 0; i < 10; i++ {
 sum += i
}
fmt.Println(sum)
```

###### FOR без итератора
```go
sum := 1
for ; sum < 1000; {
  sum += sum
}
fmt.Println(sum)
```

###### FOR без конца и края
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


### Range
```go
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

// так выведет значение (value)
for _, client := range clients {
   fmt.Println(client)
}

// так выведет оба значения(key & value)
for i, client := range clients {
   fmt.Println(i.Id)
   fmt.Println(client)
}
```

### If, switch
###### If 
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

###### Switch обычный 
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

###### Switch без условия. 
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
###### defer - выполниться после достижение и окончания последней строки ф-ции в которой она написана
```go
defer fmt.Println("world")
fmt.Println("hello")
```

###### Также можно накидывать в очередь выполнения defer (отработает по типу last-in-first-out)
```go
fmt.Println("counting")
for i := 0; i < 10; i++ {
 defer fmt.Println(i)
}
fmt.Println("done")
```

### Structs (alike Class)
###### Struct как объект
```go
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
    
    // вот тут похоже на объект из JS
    v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
```

###### Struct как класс
```go
type Vertex struct {
	name string
}

func (v Vertex) Say() {
	fmt.Println("It say Vertex method")
	fmt.Println("name =" + v.name) 
}

func main() {
	v := Vertex{name: "Krake"}  
	v.Say()
	
	gg := Vertex{name: "Vasiliy"}
	gg.Say() 
}
```
