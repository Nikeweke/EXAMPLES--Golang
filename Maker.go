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

           var choice int

           fmt.Println("======================== Hello. I am MAKER ======================= ")
           fmt.Println("type '1' to get this kit ===>[ .gitignore + pusher.bat ]")
           fmt.Println("type '2' to get this kit ===>[ .gitignore + pusher.bat + readme.md]")
           fmt.Println("type '3' to get this kit ===>[  main.go + run.bat ]")
           fmt.Println("type '4' to get this kit ===>[ .gitignore + pusher.bat + readme.md + main.go + run.bat ]")
           fmt.Println("type '8' to exit ")
           fmt.Println()
           fmt.Print("DECIDE IT NOW: ")
           fmt.Scanln(&choice)
           fmt.Println()

           switch choice {
             case 1:
                 os.Mkdir("files" , 1)
                 os.Chdir("files")
                 MakeFile(".gitignore",   "pusher.bat")
                 MakeFile("pusher.bat",    pusher)
                 fmt.Println("[ .gitignore + pusher.bat ] have been created successfuly ^^)")
             case 2:
                 os.Mkdir("files" , 1)
                 os.Chdir("files")
                 MakeFile(".gitignore",   "pusher.bat")
                 MakeFile("pusher.bat",    pusher)
                 MakeFile("readme.md",     "## Some Title")
                 fmt.Println("[ .gitignore + pusher.bat + readme.md ] have been created successfuly ^^)")
             case 3:
                 os.Mkdir("files" , 1)
                 os.Chdir("files")
                 MakeFile("run.bat",       run)
                 MakeFile("main.go",       maingo)
                 fmt.Println("[ main.go + run.bat ] have been created successfuly ^^)")
             case 4:
                 os.Mkdir("files", 1)
                 os.Chdir("files")
                 MakeFile(".gitignore",   "pusher.bat")
                 MakeFile("readme.md",     "## Some Title")
                 MakeFile("pusher.bat",    pusher)
                 MakeFile("run.bat",       run)
                 MakeFile("main.go",       maingo)
                 fmt.Println("[ .gitignore + pusher.bat + readme.md + main.go + run.bat  ] have been created successfuly ^^)")
             case 8:
                 os.Exit(3)
             default:
                 fmt.Println("OOOOHH IT IS WARNING: You typed something wrong, choose correct answer")
                 fmt.Println()
                 fmt.Println()
                 main()
           }
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


        /* ----------------------------------------------
        |  Файлы
        |
        |------------------------------------------------
        */
        // ----------------------> main.go
        var maingo = `
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
        REM go get github.com/drgomesp/gorm
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
