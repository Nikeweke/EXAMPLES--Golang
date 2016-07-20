package main


import (
       "fmt"
       "log"
      "net/http"
      "html/template"
        )


// =============================================-> MAIN FUNC <-===============================================
func main(){

  // Connect to DB
   Dbini()
   defer db.Close()

   // Для корректной работы с HTML - страничками
   http.HandleFunc("/pages/", StaticPage)

   // Что будет на странице если мы пришли по этому URL
    http.HandleFunc("/", Redirect) // Редирект на главную
    http.HandleFunc("/index/", Index)
    http.HandleFunc("/entry/", Entry)
    http.HandleFunc("/reg/",   Reg)
    http.HandleFunc("/users/", Users)

    http.HandleFunc("/show/", ShowC)
    http.HandleFunc("/del/", DelC)

    // START SERVER
    println("\nServer is RUNNING......................................."); log.Println("")
    fmt.Println(http.ListenAndServe(":5555", nil))
  }
// =============================================-> MAIN FUNC <-===============================================

// ****************************************************** INDEX
func Index(wr http.ResponseWriter, req *http.Request){

  //result, err := db.Query("INSERT INTO test(nickname, email, password) VALUES('Braker', 'Driver3@meta.ua', '12345')")
//if err != nil { fmt.Println(err,result) }
//  SetCookie(wr, "nickname", "shit")

Checker(wr, req, "pages", "index.html", "pages/templates/temp1.txt")
}
// ****************************************************** INDEX


// ****************************************************** cookas
func DelC(wr http.ResponseWriter, req *http.Request){

DelCookie(wr, "email")
DelCookie(wr, "nickname")

}

func ShowC(wr http.ResponseWriter, req *http.Request){

 email := ReadCookie(req, "email")
 password := ReadCookie(req, "password")

 fmt.Fprintf(wr,"Email: " + email + "\n")
 fmt.Fprintf(wr,"Pass: " + password + "\n")

}
// ****************************************************** cookas


// ****************************************************** Entrence
func Entry(wr http.ResponseWriter, req *http.Request){

  Checker(wr, req, "pages", "entry.html", "pages/templates/temp1.txt")
}
// ****************************************************** Entrence


// ****************************************************** Registration
func Reg(wr http.ResponseWriter, req *http.Request){

Checker(wr, req, "pages", "reg.html", "pages/templates/temp1.txt")
}
// ****************************************************** Registration


// ****************************************************** Users
func Users(wr http.ResponseWriter, req *http.Request){

Checker(wr, req, "pages", "users.html", "pages/templates/temp1.txt")


    var R [1]Mst
			    R[0] = Mst{"F1":"Степа","F2":"Котелков", "F3":"sss", "F4":"1222"}
          R[0] = Mst{"F1":"Сте22па","F2":"Котелко11в", "F3":"ss11s", "F4":"1222111"}
          R[0] = Mst{"F1":"Степ333а","F2":"Ко333телков", "F3":"s33ss", "F4":"13333222"}



set   := Mst{"Users":R}

//Checker(wr, req, "pages", "users.html", set)

templates := template.Must(template.ParseFiles("pages/index.html")) // PREPARE FILE
if err :=  templates.ExecuteTemplate(wr, "index.html", set)
err != nil { http.Error(wr, err.Error(), http.StatusInternalServerError) }

}
// ****************************************************** Users
