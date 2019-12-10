@ECHO OFF
SETLOCAL
chcp 866>nul

SET GOPATH=%CD%
SET PATH=%PATH%;%GOPATH%\BIN;

REM Linux Build settings
REM SET GOOS=linux
REM SET GOARCH=amd64
REM SET CGO_ENABLED=0

REM запускаем наш пакет(hot reloader for go)
REM fresh



REM ECHO Loading deps...
REM go get github.com/labstack/echo
REM go get github.com/sciter-sdk/go-sciter


go build -o main.exe 
.\main.exe

REM pause