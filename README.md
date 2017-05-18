# GoLang_projects

### [Пакеты для Golang](https://godoc.org/)


## Содержание

* [Что нужно для создания первого файла]()
* [Установка пакетов]()
* [Сборка под Win и Linux]()

---

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
SET GOROOT=C:\GO              rem Указываем путь где лежыт golang (куда был установлен)
SET GOPATH=%CD%               rem Указываем путь где лежат наши коды для компиляции
SET PATH=%GOROOT%\BIN;%PATH%; rem Указываем путь к Bin папке

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


  
### Сборка под Win и Linux
###### Win_build.bat
```batch
@ECHO OFF
SETLOCAL
chcp 866>nul

SET GOPATH=%CD%
SET GOROOT=C:\GO
SET PATH=%GOROOT%\BIN;%PATH%;

go build main.go
pause
```

###### Linux_build.bat
```batch
@ECHO OFF
SETLOCAL
chcp 866>nul

SET GOPATH=%CD%
SET GOROOT=C:\GO
SET PATH=%GOROOT%\BIN;%PATH%;

SET GOOS=linux
SET GOARCH=amd64
SET CGO_ENABLED=0

go build main.go
pause


```


  
  
