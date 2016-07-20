
SET GOROOT=C:\GO
SET PATH=%GOROOT%\BIN;%PATH%;
rem SET GOBIN=C:\xampp\htdocs\GO\project1\bin
SET GOPATH=C:\xampp\htdocs\GO\project1


REM ======================================== Libraries Rethink,MySQL,Socket
rem go get -u github.com/dancannon/gorethink
rem go get -u github.com/googollee/go-socket.io
rem go get -u github.com/go-sql-driver/mysql


rem go run -o main.exe
rem main.exe>>logg.txt

del main.exe
go build -o main.exe
main.exe

pause
