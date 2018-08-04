## Golang 

#### Содержание
* Install 
* Quick start
* Use external packages
* Import and cycle error
--- 

### Install
Скачать можно с офиц. сайта go - https://golang.org/. Golang сразу же установится в пути в `PATH`


### Quick start
```go
// hello.go
package main

import "fmt";

func main () {
	fmt.Println("Hello world")
	fmt.Scanln() // pause
}
```

```bash
go run hello.go
```

### Use external packages
Для того чтобы установить пакеты извне, нужно обозначить текущий рабочий проект в путях:
* **GOROOT** - это там где установлен язык. После установки, этот путь сам добавиться в `PATH`
* **GOPATH** - это путь к текущему проекту. обычно указываеться как `%CD%`. **Без неё пакеты извне будут ставиться фиг знает куда** 
* **PATH** - это глобальные переменныя windows, нам надо добавить в неё путь к бинарникам проекта, чтобы обращаться к ним просто по имени.  
    

###### packages.bat - скачка пакетов в проект
```batch
@ECHO OFF
SETLOCAL
chcp 866>nul

SET GOPATH=%CD%
SET PATH=%PATH%;%GOPATH%\BIN;

REM go get - команда которая скачивает пакеты извне
REM после одного раза можно коментить эту строку, так как пакет уже у вас

go get github.com/pilu/fresh

pause
```

###### runner.bat - запуск программы
```batch
@ECHO OFF
SETLOCAL
chcp 866>nul

SET GOPATH=%CD%
SET PATH=%PATH%;%GOPATH%\BIN;

REM for linux build
REM SET GOOS=linux
REM SET GOARCH=amd64
REM SET CGO_ENABLED=0

REM используем наш только что скачанный пакет. Это watcher за проектом.
fresh

go run main.go

pause
```

### Import and cycle error
Импорты могут замыкаться. Допустим: `main -> config -> controllers -> config`. Нельзя юзать предыдущее пространство. Надо идти дальше. Вот так: `main -> config -> controllers -> config/structs`

