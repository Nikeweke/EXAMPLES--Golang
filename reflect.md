# Reflect

#### `reflect.TypeOf(x)` 
возвращает тип переменной - значение типа `*reflect.rtype`
<br /><br />

#### `reflect.ValueOf(x)` 
возвращает ссылку на указанное значение - значение типа `reflect.Value`
<br /><br />


#### `reflect.TypeOf(x).Kind()` 
возвращает тип переменной для сравнения с reflect типами - значение типа `reflect.Kind`

```go
var x float64 = 8.6
var xType = reflect.TypeOf(x)
var xKind = reflect.TypeOf(x).Kind()

println(xType == reflect.Float64) // ERROR

println(xKind == reflect.Float64) // true
println(xKind == reflect.Array) // false
```
<br />


#### `reflect.TypeOf(x).FieldByName()` 

```go
type Contact struct {
  Name string "check: len(3, 40)"
  Id int "check: range(1, 9999)"
}

person := Contact{ "Bjork",  1234}
personType := reflect.TypeOf(person)
if nameField, ok := personType.FieldByName(""); ok {
  fmt.Printf("%q %q %q \n", nameField.Type, nameField.Name, nameField.Tag)
}
// "string" "Name" "check: len(3, 40)"

```
<br /><br />
