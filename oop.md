### OOP
Golang **DOES NOT** support enheritance only aggregation or embeding.
In golang we **CAN'T** access fields of child class(struct) in parent class(struct) methods


Aggregation and Embbeding
```go
// example 1
type Animal struct {}

type Dog struct {
  Animal // anonymus (no name) field (embbeding)
  Age int // named field (aggregation) 
}
```

```go
// example 2
import "image/color"

type ColoredPoint struct {
  color.Color // anonymus (no name) field (embbeding)
  x, y int // named field (aggregation) 
}

point := ColoredPoint{}
print(point.Color, point.x, point.y)
```


Inheritance of methods

```go
type Animal struct {
}

func (Animal) Say1() {
  println("I am Animal")
}

type Dog struct {
	Animal // embedding methods of Animal to Dog struct, now we can call "Say1" method
	name string
}

func (d *Dog) Say2(){
	println("I am Dog -", d.name)
}


func main() {
	var dog = Dog{name: "buddy"}
	dog.Say1()
	dog.Say2()
	fmt.Println(dog)
}
```

How to fire parent method with child fields
```go
// Parent struct
type BaseModel struct {}

func (this BaseModel) ToJson(model string, item interface{}) string {
	jsonData, err := json.Marshal(item )
	if err != nil {
		fmt.Println("models/" + model + "/ToJson()", err)
		return err.Error()
	}
	return string(jsonData)
}

// Child struct
type User struct {
  BaseModel
  Id            int    `json:"id"`  
  Name          string `json:"name"`
}

func (this User) ToJson() string { 
  return this.BaseModel.ToJson("User", this)
}


func main() {
	var user = User{Id: 123, Name: "buddy"}
	fmt.Println(user.ToJson())
}
```
