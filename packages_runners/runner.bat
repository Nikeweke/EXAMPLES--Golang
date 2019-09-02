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
fresh

go run main.go 
REM go build main.go
REM pause
