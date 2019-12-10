SET GOROOT=C:\GO
SET PATH=%GOROOT%\BIN;%PATH%;
 SET GOBIN=C:\xampp\htdocs\GO\project2\bin
SET GOPATH=C:\xampp\htdocs\GO\project2


REM ======================================== Libraries Rethink,MySQL,Socket
rem go get -u github.com/dancannon/gorethink
rem go get -u github.com/googollee/go-socket.io
rem go get -u github.com/go-sql-driver/mysql
rem go get -d github.com/mikespook/goemphp
rem go get -tags php7 github.com/deuill/go-php


rem go run -o main.exe
rem main.exe>>logg.txt


 del main.exe
 go build -o main.exe
 main.exe

pause
