### Interfaces

Golang **does not** support enheritance only aggregation or embeding

No `implementation` word that define type to some struct, like in other languages


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

```go
package main

import ("fmt" "math")

/* define an interface */
type Shape interface {
   area() float64
}

/* define a circle, CIRCLE not implement Shape */
type Circle struct {
   x,y,radius float64
}

/* define a rectangle Rectangle not implement Shape */
type Rectangle struct {
   width, height float64
}

/* define a method for circle (implementation of Shape.area())*/
func(circle Circle) area() float64 {
   return math.Pi * circle.radius * circle.radius
}

/* define a method for rectangle (implementation of Shape.area())*/
func(rect Rectangle) area() float64 {
   return rect.width * rect.height
}



/* define a method for shape */
func getArea(shape Shape) float64 {
   return shape.area()
}

func main() {
   circle := Circle{x:0,y:0,radius:5}
	 rectangle := Rectangle {width:10, height:5}
   
   fmt.Printf("Circle area: %f\n",getArea(circle))
   fmt.Printf("Rectangle area: %f\n",getArea(rectangle))
}
```

Another example
```go
type MyStringer interface {
	String() string
}

type Temp int

type Point struct { 
	x, y int
}

func (t Temp) String() string {
	return strconv.Itoa(int(t)) + " °C"
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}


func main() {
	var x MyStringer

	x = Temp(24) // init user-defined type
	fmt.Println(x.String()) // 24 °C

	x = Point{1, 2}
	fmt.Println(x.String()) // (1,2)
}
```
