package main

import(
        "fmt"
        "net/http"
       )



func Buttons(wr http.ResponseWriter, req *http.Request){



//=============================================== КНОПКА "ВОЙТИ"
   if req.Method == "POST" && req.FormValue("btn") == "enter" {

       // Нужно здесь : - Перехват IP
      //                - MD5 паролей
      //                - Запись в лог


  		 // получение данных с полей
  		 email    := req.FormValue("email")
  	   password := req.FormValue("password")

  		// Проверка поля на пустоту
      if !FilledOut(email) || !FilledOut(password) { fmt.Println("Заполните все поля.")
      } else{
  		        // Проверка на валидность почты
  		        if !ValidEmail(email){ fmt.Println("Адресс почты не допустим.")
              } else{
                    // Проверяем если такой пользователь
                    var user_password string // переменная для присваивание с поля "Nickname" таблицы test
                    var user_email string
                    var user_nickname string
                    db.QueryRow("SELECT email, password, nickname FROM test WHERE email = ?", email).Scan(&user_email, &user_password, &user_nickname) // Запрос на 1 результат + присваивание

                    // Проверка на существование пользователя в БД
                    if user_email == "" {  fmt.Println("Вы еще не зарегестрировались здесь.")
                     } else {
                              if user_email == email && user_password == password {


                                  fmt.Println("dddddddddddddddddddd")
                                  // Установка кук !!!!!!!!!!!!!!!!!!
                                  SetCookie(wr, "email", user_email )
                                  SetCookie(wr, "nickname", user_nickname )

                                  fmt.Println("\nEmail ->" + email)
                                  fmt.Println("Pass ->" + password)

                                } else{ fmt.Println("Неправильный пароль или логин, либо вы еще не зарегестрировались.") }
                             }
                   }
              }

     }
//=============================================== КНОПКА "ВОЙТИ"


//=============================================== КНОПКА "ЗАРЕГИСТРИРОВАТЬСЯ"
if req.Method == "POST" && req.FormValue("btn") == "reg" {

    // получение данных с полей
    email      := req.FormValue("email")
    nickname   := req.FormValue("nickname")
    repassword := req.FormValue("re-password")
    password   := req.FormValue("password")

   // Проверка поля на пустоту
   if !FilledOut(email) || !FilledOut(nickname) || !FilledOut(password) || !FilledOut(repassword) { fmt.Println("Заполните все поля.")  }

   // Проверка на валидность почты
   if !ValidEmail(email){ fmt.Println("Адресс почты не допустим.") }


   fmt.Println("\nEmail ->" + email)
   fmt.Println("Pass ->" + password)
  }
//=============================================== КНОПКА "ЗАРЕГИСТРИРОВАТЬСЯ"

}
