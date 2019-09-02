export GOPATH=$(pwd)
export PATH=$PATH:$GOPATH/bin

echo Loading packages...

go get github.com/pilu/fresh

echo Packages loaded successfuly!
