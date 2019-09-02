export GOPATH=$(pwd)
export PATH=$PATH:$GOPATH/bin

# запускаем наш пакет(hot reloader for go)
fresh

go run main.go
# go build main.go
