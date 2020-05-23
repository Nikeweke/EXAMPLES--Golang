// **FOR LOOP**. Есть массив типа - `map[User]string` , где `User` структура в которой одне поле **Id string**. Надо достать // значение которое справу (value). При обычном ренже выдает значение только ключа(key). Решение:

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

// так выведет значение справа(value)
for _, client := range clients {
   fmt.Println(client)
}

// так выведет оба значения(key & value)
for i, client := range clients {
   fmt.Println(i.Id)
   fmt.Println(client)
}
