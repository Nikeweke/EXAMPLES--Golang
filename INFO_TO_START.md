# All info to start coding Golang

### [Пакеты для Golang](https://godoc.org/)


## Содержание
* [Сборка под Win и Linux](#Сборка-под-win-и-linux)
* [Основные комманды](#Основные-комманды)
* [Что нужно для создания первого файла](#Что-нужно-для-создания-первого-файла)
* [Подключение файлов](#Подключение-файлов)
* [Установка пакетов](#Установка-пакетов)
* [Статические файлы](#Статические-файлы---js-css)
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

REM go get github.com/pilu/fresh - hot reloader
REM fresh 

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

REM go get github.com/pilu/fresh - hot reloader
REM fresh 

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

  
**3.** Создаем файл **main.go**:
```go
package main
import "fmt"

func main(){
  fmt.Println("Hello mate")
}
```
  
**4.** И компилируем нашу программу через запуск файла **run.bat**:
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

### Статические файлы - JS, CSS
```go
// при этом main.exe(main.go) лежит с папкой "public"
http.Handle("/public/",         http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
```
