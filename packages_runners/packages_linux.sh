export GOPATH=$(pwd)
export PATH=$PATH:$GOPATH/bin

echo Loading packages...

go get -u github.com/labstack/echo
echo Echo is loaded

go get github.com/fatih/color
echo Color is loaded

go get github.com/jessevdk/go-assets
echo Go-assets is loaded

echo Packages loaded successfuly!

# go get github.com/jessevdk/go-assets-builder
# go get github.com/dgrijalva/jwt-go
# go get github.com/pilu/fresh
# go get github.com/robfig/cron
# go get github.com/jessevdk/go-assets-builder

# go get github.com/jinzhu/gorm
# go get github.com/jinzhu/gorm/dialects/mysql
# go get github.com/jinzhu/gorm/dialects/sqlite
# go get github.com/jinzhu/gorm/dialects/mssql
# go get github.com/jinzhu/gorm/dialects/postgres

# go get gopkg.in/mgo.v2
