@ECHO OFF
SETLOCAL
chcp 866>nul

SET GOPATH=%CD%
SET PATH=%PATH%;%GOPATH%\BIN;

ECHO Loading packages...

go get github.com/pilu/fresh

ECHO Packages loaded successfuly!

