package main

import (
	"fmt"
  "html/template"
	"net/http"
  "log"
	 // "database/sql"
  _"github.com/go-sql-driver/mysql"
	"strings"
)





// =================================-> MAIN FUNCTION <-=======================================
func main() {

// Connect to DB
 Dbini()
 defer db.Close()

 // Что будет на странице если мы пришли по этому URL
	http.HandleFunc("/index/", Index)
	http.HandleFunc("/temp/",  Template)
	http.HandleFunc("/show/",  Show)


  // START SERVER
  println("\nServer is RUNNING...............................................\n")
  log.Println(" -> Time Started")
  fmt.Println(http.ListenAndServe(":8080", nil))
}
// =================================-> MAIN FUNCTION <-=======================================


// ========================================-> Index
func Index(wr http.ResponseWriter, req *http.Request){



  result, err := db.Query("SELECT name, surname, middle FROM test")
  if err != nil { fmt.Println(err) }


fmt.Fprintf(wr, `<html>
	              <head></head>
                 <body> `)
  i := 1
	for result.Next(){

		var name string
		var surname string
		var middle string
  	err = result.Scan( &name, &surname, &middle)
		if err != nil{ fmt.Println(err) }

    fmt.Fprintf(wr, "Запись №%v -  %s  %s %s <br>", i,name, surname, middle)
  	 i++
	}

	fmt.Fprintf(wr, `</body></html>`)


/*  if req.Method == "POST" && req.FormValue("btn") == "enter" {

		 // получение данных с полей
		 email    := req.FormValue("email")
	   password := req.FormValue("password")

		// Проверка поля на пустоту
    if !FilledOut(email) || !FilledOut(password) { fmt.Println("Email or Password is empty")  }

		// Проверка на валидность почты
		if !ValidEmail(email){ fmt.Println("Email is Shitty") }


    fmt.Println("\nEmail ->" + email)
		fmt.Println("Pass ->" + password)
   }*/

}
// ========================================-> Index



// ========================================-> Show
func Show(wr http.ResponseWriter, req *http.Request){

P:=req.URL.Path[len("/show/"):]
Z:=strings.Split(P,"*")
U:=STI(Z[0])



AA:= len(Z)
fmt.Println(AA)

if AA == 3{
	fmt.Fprintf(wr, "OKE")
}

if AA == 2{

	//fmt.Fprintf(wr,"Malo paramatrov")
  http.Redirect(wr, req, "/temp/", http.StatusSeeOther)
	return
}
V:=Z[1]
Q:=Z[2]
//Param := STI(req.URL.Path[len("/show/"):])

 var name string
 db.QueryRow("SELECT name FROM test WHERE id=?", U).Scan(&name)
 fmt.Fprintf(wr, "Name : %s %s %s",name, V, Q)


}
// ========================================-> Show


// ========================================-> Template
func Template(wr http.ResponseWriter, req *http.Request){

 var I [5]Mst
     I[0]=Mst{"S1":"ddddd22","S2":"ssssddd32"}
     I[1]=Mst{"S1":"ddddd22","S2":"ssssddd32"}
     I[2]=Mst{"S1":"ddddd33","S2":"ssssddd22"}

var R [5]Mst
			    R[0]=Mst{"Fam":"Степа","Names":"Котелков"}
      		R[1]=Mst{"Fam":"Вася","Names":"Прохоров"}
		      R[2]=Mst{"Fam":"Петя","Names":"Степанов"}
          R[3]=Mst{"Fam":"Герда","Names":"Мацареловна"}

	p :=`<table class='table'> <tr> <td>Head Office</td> </tr> </table>`


	templates := template.Must(template.ParseFiles("templates/index.html")) // PREPARE FILE

  set   := Mst{"T":p, "Y":"Наименование отчета", "U":I, "NM":R}

 if err :=  templates.ExecuteTemplate(wr, "index.html", set)
 err != nil { http.Error(wr, err.Error(), http.StatusInternalServerError) }
}
// ========================================-> Template
