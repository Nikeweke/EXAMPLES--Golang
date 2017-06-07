/**
*   Kalkulator on golang
*
*   Пример использвания функция которые возвращают несколько значений
*
*/





package main

import (
     "fmt"
     "os"
)

      /*
      |--------------------------------------------------------------------------
      |  СТАРТ
      |
      |--------------------------------------------------------------------------
      */
      func main() {

         var choice, numb1, numb2 int

         fmt.Println("|-----------------------------------------------------------------------------|")
         fmt.Println("|---->                     Welcome to Kalkulator                          <---|")
         fmt.Println("|-----------------------------------------------------------------------------|\n")

         fmt.Println("Choose desired action: ")
         fmt.Println("* Multiple -> 1")
         fmt.Println("* Sum -> 2")
         fmt.Println("* Subtraction -> 3")
         fmt.Println("* Division -> 4\n")

         fmt.Print("Your choice will be = ")
         fmt.Scan(&choice)

         fmt.Println("")
         fmt.Print("X = ")
         fmt.Scan(&numb1)

         fmt.Print("Y = ")
         fmt.Scan(&numb2)

         switch choice {
             case 1:
                 rez_m, rez_b := Multiple(numb1, numb2)
                 if rez_b { fmt.Println(rez_m) }

             case 2:
               rez_m, rez_b := Sum(numb1, numb2)
               if rez_b { fmt.Println(rez_m) }

             case 3:
               rez_m, rez_b := Subtract(numb1, numb2)
               if rez_b { fmt.Println(rez_m) }

            case 4:
               rez_m, rez_b := Division(numb1, numb2)
               if rez_b { fmt.Println(rez_m) }

            default:
                 fmt.Println("Wrong number, holly shit dude.")
         }

         AgainQuestion() // make one more operation ?
      }


      /*
      |--------------------------------------------------------------------------
      |  Умножение
      |
      |--------------------------------------------------------------------------
      */
      func Multiple(numb1, numb2 int) (string, bool){

         var rez = numb1 * numb2

         message  := fmt.Sprintf("Your result is = %d * %d = %d \n", numb1, numb2, rez)
         success  := true

        return message, success
      }



      /*
      |--------------------------------------------------------------------------
      |  Сумма
      |
      |--------------------------------------------------------------------------
      */
      func Sum(numb1, numb2 int) (string, bool){

         var rez = numb1 + numb2

         message  := fmt.Sprintf("Your result is = %d + %d = %d \n", numb1, numb2, rez)
         success  := true

        return message, success
      }



      /*
      |--------------------------------------------------------------------------
      |  Разница
      |
      |--------------------------------------------------------------------------
      */
      func Subtract(numb1, numb2 int) (string, bool){

         var rez = numb1 - numb2

         message  := fmt.Sprintf("Your result is = %d - %d = %d \n", numb1, numb2, rez)
         success  := true

        return message, success
      }



      /*
      |--------------------------------------------------------------------------
      |  Деление
      |
      |--------------------------------------------------------------------------
      */
      func Division(numb1, numb2 int) (string, bool){

         var rez = numb1 / numb2

         message  := fmt.Sprintf("Your result is = %d / %d = %d \n", numb1, numb2, rez)
         success  := true

        return message, success
      }




      /*
      |--------------------------------------------------------------------------
      |  Вопрос об еще одной операции
      |
      |--------------------------------------------------------------------------
      */
      func AgainQuestion() {

          var choice string
          fmt.Println("Еще одна операция или все на сегодня ? (y/n)")
          fmt.Scan(&choice)

          if choice == "y" || choice == "Y" {
             main()
          } else{
            os.Exit(3)
          }
      }
