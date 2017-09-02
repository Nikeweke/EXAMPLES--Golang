# Expa from Golang

### [Пакеты для Golang](https://godoc.org/)

###### run.bat
```
@ECHO OFF

REM по стандарту должна уже стоять
REM SET GOROOT=C:\GO
SET GOPATH=%CD%

REM путь к бинарникам проекта
SET PATH=%GOPATH%\BIN;%PATH%;

REM SET GOOS=linux
REM SET GOARCH=amd64
REM SET CGO_ENABLED=0


REM fresh 

go run index.go
REM go build -o build/index index.go
REM go build -o build/index.exe index.go
pause
```


  
  
