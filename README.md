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

###  Использование библиотек
Для того чтобы установить пакеты извне, нужно обозначить текущий рабочий проект в путях:
* **GOROOT** - это там где установлен язык. После установки, этот путь сам добавиться в `PATH`
* **GOPATH** - это путь к текущему проекту. обычно указываеться как `%CD%`. **Без неё пакеты извне будут ставиться фиг знает куда** 
* **PATH** - это глобальные переменныя ОС, нам надо добавить в неё путь к бинарникам проекта, чтобы обращаться к ним без полного пути.  
    
### Вопросы и Ответы

**1)** Import cycle error?
>Импорты могут замыкаться. Допустим: `main -> config -> controllers -> config`. Нельзя юзать предыдущее пространство. Надо идти дальше. Вот так: `main -> config -> controllers -> config/structs`

**2)** Где ищет `import` пакеты стандартно?
> В папке `src`. Если вы пишите: `import "storm/config"`, это значит: `./src/storm/config`

**3)** Если у меня два файла и они относяться к одному пространству(пакету). При запуске `go run main.go`, выдает ошибку когда обращяешься к функциям или переменным другого файла. Почему?
> `go run filename.go` - запускает только один файл, и поэтому не видит другие файлы из этого же пространства. 
> Проблема решаеться 2 способами:
> * I. Указать второй файл - `go run filename.go filename2.go`
> * II. Сделать билд и запустить - `go build filename.go && filename.exe`
> 

**4)** Можно ли делать десктопные приложения на Go. 
> Да, но лучше не надо. Слишком сырые инструменты для этого у Go сообщества. Лучше использовать связку Electron и Go с мостом.
