## Golang 

#### Содержание
* Установка 
* Quick start
* Использование библиотек
* Вопросы и Ответы

Примеры packages(скачивает пакеты) и runner(запускает программу):
* [packages_linux.sh](https://github.com/Nikeweke/EXAMPLES--Golang/blob/master/packages_runners/packages_linux.sh)
* [runner_linux.sh](https://github.com/Nikeweke/EXAMPLES--Golang/blob/master/packages_runners/runner_linux.sh)
* [packages.bat](https://github.com/Nikeweke/EXAMPLES--Golang/blob/master/packages_runners/packages.bat)
* [runner.bat](https://github.com/Nikeweke/EXAMPLES--Golang/blob/master/packages_runners/runner.bat)


--- 

### Установка
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

<br />


### run.bat - скрипт для запуска 

```batch
@ECHO OFF
SETLOCAL
chcp 866>nul
CLS

SET BUILD_PATH=.\myapp

REM you can specify for downloading packages to project folder
REM SET GOMODCACHE=%CD%\packages

REM download all packages that project requires
go mod tidy

ECHO Building service...

rem build with flags
go build -o %BUILD_PATH%.exe

rem start in dev mode (default)
%BUILD_PATH%.exe
```

<br />

### GOPATH
Эта системная переменная которая указывает куда скачивать пакеты нужные для проектов.

### GOMODCACHE
Если установить при сборке эту переменную то можно проигнорировать `GOPATH` и брать пакеты из указанной папки. Допустим для проекта нужны определенные пакеты, а не брать из общего списка `GOPATH`. Пример: `SET GOMODCACHE=%CD%\packages`

<br />

### Clear cache

###### очистит все пакеты по пути GOPATH

```
go clear -modcache
```

---

### Вопросы и Ответы

**1)** Import cycle error?
>Импорты могут замыкаться. Допустим: `main -> config -> controllers -> config`. Нельзя юзать предыдущее пространство. Надо идти дальше. Вот так: `main -> config -> controllers -> config/structs`

**2)** Если у меня два файла и они относяться к одному пространству(пакету). При запуске `go run main.go`, выдает ошибку когда обращяешься к функциям или переменным другого файла. Почему?
> `go run filename.go` - запускает только один файл, и поэтому не видит другие файлы из этого же пространства. 
> Проблема решаеться 2 способами:
> * I. Указать второй файл - `go run filename.go filename2.go`
> * II. Сделать билд и запустить - `go build filename.go && filename.exe`
> 

**3)** Можно ли делать десктопные приложения на Go?
> Да, но лучше не надо. Слишком сырые инструменты для этого у Go сообщества. Лучше использовать связку Electron и Go.
> (Испробовано: zserge/webview, wails, gtk/glade, sciter, Qt, go-astielectron, WebAssembly)
