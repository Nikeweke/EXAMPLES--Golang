### Structs



```go
type Student struct {
    Name string
    Age  int
}
```

3 ways to init struct variable

```go
// 1. 
var s Student // same as - Student{"", 0}
```

```go
// 2. 
// The `new` keyword can be used to create a new struct. It returns a pointer to the newly created struct.
var s2 = new(Student) // same as - &Student{"", 0}
```

```go
// 3. 
var s3 = Student{} // same as - Student{"", 0}
```

Inheritance of methods

```go
type Animal struct {
}

func (Animal) Say() {
	println("I am Animal")
}

type Dog struct {
	Animal // embedding methods of Animal to Dog struct, now we can call "Say" method
	name string
}

func (d *Dog) Say() {
	println("I am Dog -", d.name)
}


func main() {
	var dog = Dog{name: "buddy"}

	dog.Say()
	fmt.Println(dog)
}


```
