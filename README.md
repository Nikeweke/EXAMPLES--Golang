# Expa from Golang

### [Пакеты для Golang](https://godoc.org/)

* Websockets - `go get github.com/gorilla/websocket`
* Jobs - `go get github.com/robfig/cron`
* Gorm(ORM for DB) - `go get github.com/drgomesp/gorm`
* Mysql driver - `go get github.com/go-sql-driver/mysql`
* Hot reloader - `go get github.com/pilu/fresh`

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

REM go get github.com/pilu/fresh
REM fresh 

go run index.go
REM go build -o build/index index.go
REM go build -o build/index.exe index.go
pause
```


  
  
