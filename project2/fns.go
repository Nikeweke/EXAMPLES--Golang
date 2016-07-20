
package main

import (
        "regexp"
        "net/http"
        "log"
         "database/sql"
        _"github.com/go-sql-driver/mysql"
        "html/template"
        "path"
        "time"
        "fmt"

       )

// INTERFACE for out from DB
type Mst map[string]interface{}



// ---------------------------------------------------------------------------- REDIREDKTER
func Redirect(wr http.ResponseWriter, req *http.Request){

  Index(wr, req)

}
// ---------------------------------------------------------------------------- REDIREDKTER

// ---------------------------------------------------------------------------- STATIC-PAGER
func StaticPage(wr http.ResponseWriter, req *http.Request){
     // Allows
    // w.Header().Set("Access-Control-Allow-Origin",  "*")

	   //  File static page
   	http.ServeFile(wr, req, req.URL.Path[1:])
}
// ---------------------------------------------------------------------------- STATIC-PAGER






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
func SetCookie(wr http.ResponseWriter, NameCookie, ValueCookie string){
fmt.Println("Name ->" + NameCookie)
fmt.Println("Val ->" + ValueCookie)
  t := 86400*5 // Куки на 5 дней
	cookie := &http.Cookie{ Name: NameCookie,
                          Value: ValueCookie,
                          HttpOnly: true,
                          MaxAge: t,
                          Path: "/"  }
	http.SetCookie(wr, cookie)
}
//==================================


/********************************************************************************************************************************
 *  TITLE         : Запись кук Write Cookie Days на период в днях
 *  DATE          : 17-04-2015 15:03
 *  DESCRIPTION   : NameCookie    - имя
 *                  ValueCookie   - значение стринг
 *                  Hrs           - Период в часах
 *  USAGE         : Sys_Wch(w, "Primer","Test",8)
 *********************************************************************************************************************************/
func Sys_Wcd(wr http.ResponseWriter, NameCookie, ValueCookie string, Hrs time.Duration) {
	cookie := &http.Cookie{Name: NameCookie, Value: ValueCookie, Expires: time.Now().Add(356 * Hrs * time.Hour), HttpOnly: true,  Path: "/"}
	http.SetCookie(wr, cookie)
}







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


//////////////////////////////////////////////////////////////////////////////////////// PREPARE PAGES-HTML FOR HTML-TEMPLATES(METAH, HEAD, FOOTER)
// Эта функция нужна для того что бы подключать html-шаблоны, такие как : METAH, HEAD, FOOTER.
// Без этой функции - будут кракозябры на странице-html отображены
func Checker(wr http.ResponseWriter, req *http.Request, folder, page, temp string){

  // Parsing
  tmpl, err := template.ParseFiles(path.Join(folder, page), temp)
  if err != nil { http.Error(wr, err.Error(), http.StatusInternalServerError);  return } // Error
  if err := tmpl.Execute(wr, nil); err != nil { http.Error(wr, err.Error(), http.StatusInternalServerError)  }

  // Проверка на нажатие кнопок для обработки действий - core.go
  Buttons(wr, req)
}
//////////////////////////////////////////////////////////////////////////////////////// PREPARE PAGES-HTML FOR HTML-TEMPLATES(METAH, HEAD, FOOTER)
