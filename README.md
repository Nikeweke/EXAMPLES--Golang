# Expa from Golang

### [Пакеты для Golang](https://godoc.org/)


## Содержание
* [Сборка под Win и Linux](#Сборка-под-win-и-linux)
* [Основные комманды](#Основные-комманды)
* [Что нужно для создания первого файла](#Что-нужно-для-создания-первого-файла)
* [Подключение файлов](#Подключение-файлов)
* [Установка пакетов](#Установка-пакетов)
* [Проблемы и их решение](#Проблемы-и-их-решение)

---

### Сборка под Win и Linux
###### Win_build.bat
```batch
@ECHO OFF
SETLOCAL
chcp 866>nul

REM по стандарту должна уже стоять
REM SET GOROOT=C:\GO
SET GOPATH=%CD%

REM путь к бинарникам проекта
SET PATH=%GOPATH%\BIN;%PATH%;
REM SET PATH=%GOROOT%\BIN;%PATH%;

go run main.go
pause
```

###### Linux_build.bat
```batch
@ECHO OFF
SETLOCAL
chcp 866>nul

SET GOROOT=C:\GO
SET GOPATH=%CD%
SET PATH=%GOROOT%\BIN;%PATH%;

SET GOOS=linux
SET GOARCH=amd64
SET CGO_ENABLED=0

go build main.go
pause


```


### Основные комманды
* **go env** - можно посмотреть установлены ли переменные окружения, GOPATH GOROOT
* **go run main.go** - запускает, но не компилирует
* **go build main.go** - компилит проект в корень
* **go build -o bin/app.exe main.go** - компилит проект в ./bin

### Что нужно для создания первого файла:
**1.** Установить язык
```
https://golang.org/dl/
```
**2.** В проекте сделать файл который будет компилить все (**run.bat**):
```go
@ECHO OFF
SETLOCAL
chcp 866>nul

rem Установка переменных среды для компиляции
rem Указываем путь где лежыт golang (куда был установлен) 
SET GOROOT=C:\GO
rem Указываем путь где лежат наши коды для компиляции
SET GOPATH=%CD%
rem Указываем путь к Bin папке
SET PATH=%GOROOT%\BIN;%PATH%;

go run app/main.go   rem запустит но не выплюнет закомпиленный файл
rem go build -o bin/app.exe app/main.go app/fns.go    rem Скомпилирует и положит в bin/

rem Если 2 файла
rem go run app\main.go app\fns.go
```

**3.** Создаем в папке проекта: 3 директории
  + *src*
  + *pkg*
  + *bin*
  
**4.** Создаем файл **main.go**:
```go
package main
import "fmt"

func main(){
  fmt.Println("Hello mate")
}
```
  
**5.** И компилируем нашу программу через запуск файла **run.bat**:
```
start run.bat
```


### Подключение файлов
Структура такая:
* run.bat
* main.go
* pkg
* bin
* **src**
   + app
        - fns / fns.go (пространство имен - **package fns**, иначе будет ошибка, все связано по папкам) 




###### run.bat
```batch
@ECHO OFF
SETLOCAL
chcp 866>nul

rem Установка переменных среды для компиляции
SET GOROOT=C:\GO             
SET GOPATH=%CD%               
SET PATH=%GOROOT%\BIN;%PATH%; 

go run main.go   rem запустит но не выплюнет закомпиленный файл

rem Если 2 файла
rem go run app\main.go app\fns.go
```

###### main.go
```go
package main

import (
    "app/fns"   // <!-- мой пакет лежит здесь : корень / src/app/fns/fns.go (где )
    "fmt"
)

func main(){
   fmt.Println("Hello humano")
   
   fns.Braker()  // <! -- это функция из моего пакета, ф-ция должна быть с Большой буквы, 
                 //       иначе даже после импорта не сможешь её заюзать
   
   fmt.Scanln()
}
```

###### src/app/fns/index.go
```go
package fns

import "fmt"

func Braker(){
    fmt.Println("Hello its braker")
}
```



### Установка пакетов
Установка пакетов делаеться через файл **run.bat**:
```batch
@ECHO OFF
SETLOCAL
chcp 866>nul


rem Установка переменных среды для компиляции
SET GOROOT=C:\GO
SET GOPATH=%CD%
SET PATH=%GOROOT%\BIN;%PATH%;

go get github.com/therecipe/qt
pause
```


  
### Проблемы и их решение

1. **FOR LOOP**. Есть массив типа - `map[User]string` , где `User` структура в которой одне поле **Id string**. Надо достать значение которое справу (value). При обычном ренже выдает значение только ключа(key). Решение:
```golang

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
```

2. **CHANNELS**. Каналы работают только в goroutines (потоках) иначае ошибка - deadlock. 
```golang
package main
import "fmt"

//MAIN
func main() {
    c := make(chan int, 100)
    
    // Вот так не будет работать
    c <- 1
    c <- 2
    numb := <-c
    fmt.Prinln(numb) // deadlock !
    
    // Вот так работает
    go func(){
      c <- 2
      c <- 1
      c <- 4
      c <- 5
   }()

    fmt.Println(<-c)
    fmt.Println("THAT LEN",  len(c))

    fmt.Println(<-c)
    fmt.Println("THAT LEN",len(c))
    fmt.Println(<-c)
    fmt.Println(<-c)
    
    fmt.Scanln()
}

```

  
  
