
package main

import (
        "regexp"
        "net/http"
        "log"
         "database/sql"
        _"github.com/go-sql-driver/mysql"
        "strconv"
       )

// INTERFACE for out from DB
type Mst map[string]interface{}



func STI(StrParameter string) int64 {
	k, err := strconv.ParseInt(StrParameter, 10, 64)
	if err != nil {
		k = 0
	}
	return k
}


////////////////////////////////////////////////////////////////////////////////////////// CONNECT TO DB
//================================== Действие: Connect to DB, thru var - db
var db *sql.DB
func Dbini(){
	 dbs, err:= sql.Open("mysql", "root:@tcp(localhost:3306)/golang")
	 if err != nil{ log.Println(err) }
	 db=dbs
}
//==================================
//////////////////////////////////////////////////////////////////////////////////////////  CONNECT TO DB



////////////////////////////////////////////////////////////////////////////////////////// AUTH
//================================== Действие: Проверяет заполены ли поля
func FilledOut(field string) bool{

  if field == "" { return false
   } else{ return true }
}
//==================================

//================================== Действие: Проверяет валидность почты
func ValidEmail(email string) bool{

  symbols := regexp.MustCompile(`^[a-zA-Z0-9_\.\-]+@[a-zA-Z0-9\-]+\.[a-zA=z0-9\-\.]+$`)

   if symbols.MatchString(email){ return true
   } else{ return false }
}
//==================================
////////////////////////////////////////////////////////////////////////////////////////// AUTH




//////////////////////////////////////////////////////////////////////////////////////// COOKIES
//================================== Действие: Ставит куки
func SetCookie(wr http.ResponseWriter, Obj, Obj1 string){

  t := 86400*5 // Куки на 5 дней
	cookie := &http.Cookie{ Name: Obj,
                          Value: Obj1,
                          HttpOnly: false,
                          MaxAge: t,
                          Path: "/"  }
	http.SetCookie(wr, cookie)
}
//==================================


//================================== Действие: Удаляет куки
func DelCookie(wr http.ResponseWriter, NameCookie string) {
   t := -86400*100 // Куки на 5 дней
	 cookie := &http.Cookie{ Name: NameCookie,
                           Value: "",
                           MaxAge: t,
                           Path: "/"    }
	 http.SetCookie(wr, cookie)
}
//==================================


//==================================
func ReadCookie(req *http.Request, NameCookie string) string {
	ck, err := req.Cookie(NameCookie)

	// Error
	if err != nil {
	   return ""
	}  else {
	   return string(ck.Value)
	}
}
//==================================
//////////////////////////////////////////////////////////////////////////////////////// COOKIES
