### OOP

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
