# Type cast cases

### 1. 
###### Неправильно
```go
// Ситуация когда есть функция с параметрами типа `func SomeFn(data interface{}) (interface{}, error)`
func SomeFn(data interface{}) (interface{}, error) {
  // ... тут происходит допустим запрос и мы получаем bytes []byte
  
  // мы хотим загнать байты в структуру
  var err := json.Marshal(bytes, &data)
  return data, nil
}

func main() {
  type User struct {}
    
  // тут мы получим пустую переменную типа interface{}
  var user, err = SomeFn(User{})
}
```

###### Правильно
```go
// 1. переписать "SomeFn"
func SomeFn(data interface{}) error {
```

```go
// 2. передавать переменную по указателю
func main() {
  type User struct {}
  var user = User{}
  var err = SomeFn(&user)
}
```
