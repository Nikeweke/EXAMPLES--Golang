package main

import (
   "os"
   "fmt"
)

        /* ----------------------------------------------
        |  Старт
        |
        |------------------------------------------------
        */
        func main(){


// ----------------------> main.go
           var main = `
package main
import "fmt"

func main(){
 fmt.Println("Hello its me, Mario")
}
`

// ----------------------> run.bat
           var run = `
@ECHO OFF

REM по стандарту должна уже стоять
REM SET GOROOT=C:\GO
SET GOPATH=%CD%

REM путь к бинарникам проекта
SET PATH=%GOPATH%\BIN;%PATH%;

REM SET GOOS=linux
REM SET GOARCH=amd64
REM SET CGO_ENABLED=0

REM go get github.com/gorilla/websocket
REM go get github.com/robfig/cron
REM go get -u github.com/jinzhu/gorm
REM go get github.com/go-sql-driver/mysql

go run main.go
REM go build main.go
pause
`


// ----------------------> pusher.bat
           var pusher = `
git init
git remote add origin https://github.com/.....
git add -A
git commit -m "new coomit"
git push -u origin master
REM git push
`


           // File making
           MakeFile(".gitignore",   "pusher.bat")
           MakeFile("readme.md",     "## Some Title")
           MakeFile("pusher.bat",    pusher)
           MakeFile("run.bat",       run)
           MakeFile("main.go",       main)
        }



        /* ----------------------------------------------
        |  Создание файла
        |
        |------------------------------------------------
        */
        func MakeFile(name, text string){
            file, err := os.Create(name)
            checkErr(err)
            defer file.Close()
            file.WriteString(text)
        }




        /* ----------------------------------------------
        |  Решала ошибок
        |
        |------------------------------------------------
        */
        func checkErr(err error){
          if err != nil {
              fmt.Println(err)
              // os.Exit(1)
          }
        }
