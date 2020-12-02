# Array to object
Makes from array - `[{id, name}, ...]` this - `{id:{id,name}, ...}`

```go
func ArrayToObj(field string, arr interface{}) map[string]interface{} {
	var result = map[string]interface{}{}
	var arrValue  = reflect.ValueOf(arr) // convert to reflect.Value
	var arrFirstValField = arrValue.Index(0).FieldByName(field)

	if arrValue.Len() == 0 {
		return result
	}

	if arrFirstValField.Kind() == reflect.String {
		for i := 0; i < arrValue.Len(); i++ {
			var item  = arrValue.Index(i).Interface()
			var id = arrValue.Index(i).FieldByName(field).Interface().(string)
			result[id] =  item
		}
	} 

	if arrFirstValField.Kind() == reflect.Int {
		for i := 0; i < arrValue.Len(); i++ {
			var item  = arrValue.Index(i).Interface()
			id   := strconv.Itoa(arrValue.Index(i).FieldByName(field).Interface().(int) )
			result[id] =  item
		}
	}

	if arrFirstValField.Kind() == reflect.Int64 {
		for i := 0; i < arrValue.Len(); i++ {
			var item  = arrValue.Index(i).Interface()
			var idn   = arrValue.Index(i).FieldByName(field).Interface().(int64)
			id        :=  strconv.FormatInt(idn, 10)
			result[id] =  item
		}
	}

	return result
}
```
