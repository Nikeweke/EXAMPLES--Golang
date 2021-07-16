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
    
  // тут мы получим пустую переменную типа interface{}, а не User как ожидалось
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

### 2. Switch type 

Если мы хотим для определенных данных сделать какой-то вызов или дополнить можно юзать `switch type` помогает сразу
закастовать переменную и работать с ней.

```go
	switch v := items.(type) {
   	case []models.Product:
			v = services.AddProviderSearchResults(filters.Search, v)
			return c.JSON(http.StatusOK, v)     
		default:
      return c.JSON(http.StatusOK, items)
	}
```
