 

**********************************************************************

                    ### Особенности GO от С++  #### 

**********************************************************************

Отличие от С++ , ГО
-----------------------


   package main  ==   #include<iostream> 
   
 

  import (             ////// =  using namespace std

        "fmt"   ///// -- подключение биоблиотеки ввода и вывода   
  )



func main() {   /////= int main() {}

      var n int    ///// int n 
      fmt.Println(n)   ///// cout << n;

    
    in := ""          ///// Вместо -> system(pause)
    fmt.Scanln(&in)


   //Либо так можно сделать в конце паузу на 10 сек. 

  //duration := time.Duration(10)*time.Second
 // time.Sleep(duration)
----------------------------------------------------------------------






                         
**********************************************************************

                    ###  ORIGINS of GO   #### 

**********************************************************************
         
-------------------------------------------------------

**Присваивание переменным значений : 
          
          x := "Hello"  или так  var x = "Hello"
         
          x := 5 или так var x = 5
-------------------------------------------------------



-------------------------------------------------------

**Обьява переменной 
          
           var x int
           var s string
           var d int32 
           var b float32
           var q float64
-------------------------------------------------------



-------------------------------------------------------
**Вывод


         x := "dog"

         fmt.Println("HE is ", x)

-------------------------------------------------------



-------------------------------------------------------
**Цыкл

var i int 
	
	for i =0; i <= 10; i++  {
		
		fmt.Println(i)		
	}
-------------------------------------------------------






-------------------------------------------------------
**Цикл(FOR) + IF

for i := 1; i <= 10; i++ {
           
		       if i % 2 == 0 { fmt.Println(i, "четно")     ///// вот это правильный вид IF 

			         }  else { fmt.Println(i, "нечетно")} 
           
		 }

                     if i % 2 == 0 { fmt.Println(i, "четно") **}**     ///// вот это нерабочий вид IF 
                         else { fmt.Println(i, "нечетно")}            //// все из за " } " , если нет 
                                                                  //// дужки перед ELSE не работать   


-------------------------------------------------------





-------------------------------------------------------

** что бы окно консоли не закрывалось 

import ( "time"
         )

   // Закрытие через 10 мни
   duration := time.Duration(10)*time.Second
   time.Sleep(duration)

-------------------------------------------------------





------------------------------------------------------- 
**Считивание # SCANF #

fmt.Print("Enter:")
var c int
fmt.Scanf("%d", &c)
fmt.Print(c)

-------------------------------------------------------



------------------------------------------------------- 
** Закрытие по нажатию кнопки Enter

import ( "time"
            "fmt"
               "bufio"
                  "os"


         )

fmt.Print("Print Enter to continue....")
bufio.NewReader(os.Stdin).ReadBytes('\n')

-------------------------------------------------------





-------------------------------------------------------  
** Переобразование типа данных

var total float64
len(x) -- int

fmt.Print(total/ float64(len(x)))
//  
-------------------------------------------------------





-------------------------------------------------------
**Обьява массива

x := [5]float64{98, 93, 23, 22, 22}

или так

x := [5]float64{

          98,
          93,
          23,
          22,
          21,
      }

-------------------------------------------------------




-------------------------------------------------------
** Massive float + цикл проход по масиву без итератора  


var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	
	var total float64 = 0
	for _, value:= range x {
		
			
		total = total + value
		
	}
	
	fmt.Println(total/float64(len(x)))	

-------------------------------------------------------




-------------------------------------------------------
**Срезы

// создание среза x := make([]float64, 5)

slice1 := []int{1,2,3}
slice2 := make([]int, 2)
copy(slice2, slice1)
fmt.Println(slice1, slice2)

-------------------------------------------------------





-------------------------------------------------------
**Создание карты(map) типа string - string

elements := make(map[sting]string)

    elements["H"] = "Hydrogen"
	elements["He"] = "Hellium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Berrylium"

	fmt.Println(elements["H"])

-------------------------------------------------------




--------------------------------------------------
**Еще один способо обьявы элементво в карте

elements := map[string]string{
"H": "Hydrogen",
"He": "Helium",
"Li": "Lithium",
"Be": "Beryllium",
"B": "Boron",
"C": "Carbon",
"N": "Nitrogen",
"O": "Oxygen",
"F": "Fluorine",
"Ne": "Neon",
}

--------------------------------------------------





-------------------------------------------------
** Проверка на существование элемента в карте(map)

elements := make(map[string]string)
	
	elements["H"] = "Hydrogen"
	elements["He"] = "Hellium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Berrylium"
	
	
	name, ok := elements["He"]
	
	fmt.Println(name, ok)

-------------------------------------------------




-------------------------------------------------
// MAP Mendels
	elements := map[string]map[string]string{
		
		 
		         "H": map[string]string{
					"name" : "Hydrogen",
					"state" : "gas",
				  },
				
				
				"He" : map[string]string{
					   "name" : "Hellium",
					   "state" : "gas" , 
					 },
				
				
				"Li" : map[string]string{
					  "name" : "Lithium",
					   "state" : "solid",								
				    },	
		
	}
	
	
	
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
		
	}
-------------------------------------------------	





*************************************************
                 ### FUNCTIONS ###
*************************************************

-------------------------------------------------
** Возраврат значений функцией 

func f() (int, int){       // Ф-ция "f" на входе int

               return 5,6 // Вертае 5 и 6 в ф-цию "main", x & y
           }


func main() {

           x, y := f()
           fmt.Print(x,y)
         }
-------------------------------------------------




-------------------------------------------------
** Динамическое кол-во переданых атрибутов в ф-цию

func add(wgb ...int) int { // ф-ция "add" принимает значения 
	                       // int в переменную "wgb" .Вертае int
            
            total := 0     // Складивает в себя все переменные что прошли
                           //  через цикл  

            for _, v := range wgb { total += v }  // этот цикл

            return total   // вовзрат суммы        
}


func main() {

      fmt.Println( add ( 1 , 2 , 3) )
                              // Отправка в ф-цию "add" чисел 1, 2, 3
                              // и возваращает сумму , то есть total 


      /*
         // еще можно так с срезом сделать, как пример:

            xs := []int{1,2,3} // Срез "xs" с целыми числами 1,2,3

            fmt.Println( add (xs...) ) //  Отправка среза "xs" и 
                                       //  и возврат той же суммы "total"
      */
  }

-------------------------------------------------






-------------------------------------------------
** Ф-ция внутри главной ф-ции

func main() {
      
      add := func ( x, y int) int {return   x + y } //Ф-ция "add"  
                                                    //input -- 2 inta(x ,y); 
                                                    //out -- int (x + y)


      fmt.Println( add (1,1) ) // Вывод результата который вылетел из
                               // ф-ции "add", в которую полетели 2 инта
                               // и вернулся инт
}

-------------------------------------------------





-------------------------------------------------
**Замыкание , фиг пойми как работать
/*

Один из способов использования замыкания — функция, возвращающая другую функцию, которая при вызове
генерирует некую последовательность чисел. Например, следующим образом мы могли бы сгенерировать все
четные числа:


*/

func makeEvenGen() func() uint {
	
	i := uint(0)
	return func() (ret uint){
	
	ret = i
	i = i + 2
	return
     }

 }

func main() {

nextEven := makeEvenGen()

fmt.Println(nextEven()) // 0
fmt.Println(nextEven()) // 0
fmt.Println(nextEven()) // 0
}




-------------------------------------------------





-------------------------------------------------
** Рекурстя неясная ???

func fact(x uint) uint {
	
	if x == 0 { return 1 }
	
	return   x * fact( x - 1 )
	
}


func main() {
	
	
	fmt.Println(fact(4))
}

-------------------------------------------------


*************************************************
                 ### Указатели ###
*************************************************

-------------------------------------------------
** /*В Go указатели представлены через оператор * (звёздочка),
 за которым следует тип хранимого значения. 
 В функции "zero" "prom" является указателем на int .

* также используется для «разыменовывания» указателей. 
Когда мы пишем *prom = 0 , то читаем это так: «храним
int 0 в памяти, на которую указывает prom ». 
Если вместо этого мы попробуем написать prom = 0 то получим
ошибку компиляции, потому что prom имеет тип не int , 
а *int . Соответственно, ему может быть присвоен только
другой *int .

Также существует оператор & , который используется для получения адреса переменной.
 &x вернет *int (указатель на int ) потому что x имеет тип int . 
 Теперь мы можем изменять оригинальную переменную. 
 &x в функции main и prom в функции zero указывают на один и тот же участок в памяти.
 */

// Тут указатель нужен для того что бы поменять значение переменной
  // в главной функции
func zero(prom *int){

     prom = 0
}


func main() {

  x := 5  // "x" int 

  zero ( &x )       // Отправка в ф-цию "zero" -> "x"
  fmt.Println(x)   // Out of changed "x"


}
-------------------------------------------------


************************************************
          ### Struct and interface ###
*************************************************

-------------------------------------------------
package main 

import ( "fmt"; "math"  )


// ** CIRCLE **
type Circle struct { // Типа как класс 
	
	x , y, r float64	// эть поля float64
}



// ** Rectangle **
type Rectangle struct {
	
	x1, y1 float64	
	x2, y2 float64	
}



//FUNC for rEctange
func (r *Rectangle) area() float64 {	
	l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w	
}


// Если бы "(c *Circle) " была без "*" тогда 
// не изменятся наши переменные после действия
//в ф-ции "circleArea"

//FUNC for Circle
func (c *Circle) area() float64 {     // area() - название ф-ции; float - out data
                                      // (c *Circle) - "получатель"

          return math.Pi * c.r*c.r
}




func main() {	
		
	c := Circle{0, 0, 5} // Екземпляр нашего класса Circle и передача туда значений
	                     // или так : "c := Circle{x: 0, y: 0, r: 5}"
	
	fmt.Println( с.area() )	                 
 	
 	r := Rectangle{ 0, 0, 10, 10}
	fmt.Println( r.area() )	
}
-------------------------------------------------





************************************************
          ### GO ROUTINE & CHANNELS ###
*************************************************


-------------------------------------------------
** example of GO ROUTINE
package main 

import "fmt"

// FUNC f
func f (n int) {
	
	for i := 0; i < 10; i++{ fmt.Println(n, ":", i)  }	
}


func main() {
	
	//LAUNCH GO ROUTINE
	go f(0)
	
	//WITHOUT THIS STUFF IT doesnt work
	input := "" // or @var input string@ 
	fmt.Scanln(&input)
}
-------------------------------------------------



-------------------------------------------------
** go routine z zadershka
package main

import ("fmt"; "time"; "math/rand")



func f(n int) {
	for i:= 0; i<10 ; i++ {
		
		fmt.Println(n,":", i)
		
		//time sleep zadershka
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
		}	
}



func main() {
	
	for i:=0; i<10; i++ {go f(i)}
	
	in := ""
	fmt.Scanln(&in)
	
}
-------------------------------------------------


-------------------------------------------------
** WORK OF CHANNELS
package main
import ("fmt"; "time")


//PINGER
func pinger(c chan string) { for i := 0; ; i++ { c <- "ping" } }


//PONGER
func ponger(c chan string) { for i:=0 ; ; i++ { c <- "pong" }}



//OUTTER
func printer(c chan string) {
        for {
              msg := <- c
              fmt.Println(msg)
              time.Sleep(time.Second * 1)
            }
  }

//MAIN
func main() {
     var c chan string = make(chan string) // CREATE c example for using functions 

//ROUTINES
//Вывод "ping pong" по переменно в ф-цию "printer" доносится переменная "с" с её 
// значением . 
go ponger(c)
go pinger(c)
go printer(c)


//ENDs 
in := ""
fmt.Scanln(&in)
}
-------------------------------------------------







-------------------------------------------------
**CHANGING CHANNELS OR @SELECT@ FOR CHANNELS
package main

import ("fmt"; "time" )

// MAIN
func main(){
	
	//CREATE CHANNELS
	c1 := make(chan string)
	c2 := make(chan string)
	
	//FUNC 1 that send in c1 -> "from 1"
	go func(){
		for {
			c1 <- "from 1"
			time.Sleep(time.Second*2)	
			
		}
	}()// <- !!!
	
	
	//FUNC 2 that send in c2 -> "from 2"
	go func() {
		
		for{
			c2 <- "from 2"
			time.Sleep(time.Second*3)
					
		}		
	}()
	
	
	//FUNC 3 that takes signals that made FUNC 1-2 
	//from CHANNELS "c1" & "c2". SELECt, takes dicision about what "msg" will print 
	go func(){
		for{
			select {
				case msg1 := <- c1:
				    fmt.Println(msg1)
				case msg2 := <- c2:
				    fmt.Println(msg2)				
			   }		
			
		}	
		
	}()
	
	//END
	var input string 
	fmt.Scanln(&input)

/*Когда готовы несколько каналов, получение сообщения происходит из случайно выбранного готового канала. 
Если же ни один из каналов не готов, оператор блокирует ход программы до тех пор, 
пока какой-либо из каналов будет готов к отправке или получению.*/	
}
-------------------------------------------------





************************************************
          ### WORK WITH PACKET @STRING@ ###
*************************************************


-------------------------------------------------
package main
import ("fmt"; "strings")

func main() {
     
      fmt.Println(
				
				// true
				strings.Contains("test", "es"),
				
				// 2
				strings.Count("test", "t"),
				
				// true
				strings.HasPrefix("test", "te"),
				
				// true
				strings.HasSuffix("test", "st"),
				
				// 1
				strings.Index("test", "e"),
				
				// "a-b"
				strings.Join([]string{"a","b"}, "-"),
				
				// == "aaaaa"
				strings.Repeat("a", 5),
				
				// "bbaa"
				strings.Replace("aaaa", "a", "b", 2),
				
				// []string{"a","b","c","d","e"}
				strings.Split("a-b-c-d-e", "-"),
				
				// "test"
				strings.ToLower("TEST"),
				
				// "TEST"
			      strings.ToUpper("test"),
                )
}
-------------------------------------------------




************************************************
          ### READ and WRITE FILES###
*************************************************


-------------------------------------------------
**READ FILE.txt
________________
package main

import("fmt"; "io/ioutil")

func main(){
	
	bs, err := ioutil.ReadFile("C:/test.txt")
	if err != nil{ return }
	
	str := string(bs)
	fmt.Println(str)	
}
-------------------------------------------------


-------------------------------------------------
**WRITE file


package main

import( "os")

func main(){
	//CREATE
	file, err := os.Create("test.txt")
	if err != nil{ 
	
	//handle error here
	return }
	
	defer file.Close()
	
	file.WriteString("test") // That will be in test.txt
	
}
-------------------------------------------------


************************************************
          ### SERVERS###
*************************************************

-------------------------------------------------
**TCP SERVER
package main

import (
	"encoding/gob"
	"fmt"
	"net"	
)


//FUNC-SERVER
func server(){
	
	//listen on a port
	ln, err := net.Listen("tcp", ":9999")
	   if err != nil {
		        fmt.Println(err)
				return		
	     }
		
		//for
		for{
			//accept a conection
			c, err := ln.Accept()
			if err != nil{
			        fmt.Println(err)		
				    continue
			}
			
			//handle connection
	         go handleServerConnection(c)
		}
	
}


//FUNC HANDLING SERVER
func handleServerConnection(c net.Conn) {
	
	//receive msg
	var msg string
	
	err := gob.NewDecoder(c).Decode(&msg)
	
	if err != nil {fmt.Println(err)
 	     } else {fmt.Println("Received",msg)}
	
	c.Close()	
}




//FUNC CLIENT
func client(){
	//connetct to serv
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
			}

   //send msg
      msg := "ZEig Heil"
	 fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {fmt.Println(err)}
	
	
	
	c.Close()		
}

func main(){
	
	go server()
	go client()
	
	in := ""
	fmt.Scanln(&in)	
}
-------------------------------------------------





-------------------------------------------------
**HTTP SERVER
// To see work it or not , just write in brouzer @ http://127.0.0.1:9000/hello @ 
package main
import ("net/http" ; "io")

func hello(res http.ResponseWriter, req *http.Request) {
          res.Header().Set(
                            "Content-Type",
                             "text/html",
                                         )

          io.WriteString(
                          res,
                      `<doctype html>
					
                       <html>
 						<head>
							<title>Heild Gidra</title>
						</head>
						
						<body>
						Hello World!
						</body>
						
						</html>`,
                     )
  }



func main() {
http.HandleFunc("/hello", hello)
http.ListenAndServe(":9000", nil)

http.Handle(
         "/assets/",
           http.StripPrefix(
                  "/assets/",
                    http.FileServer(http.Dir("assets")),
                                      ),
                        )
}

-------------------------------------------------











************************************************
          ### OTHER STUFF###
*************************************************


-------------------------------------------------
**OWN ERRORS
package main
import "errors"

func main() {
err := errors.New("error message")
}
-------------------------------------------------


-------------------------------------------------
**LIST
package main
import("fmt"; "container/list")

func main(){
     var x list.list
     x.PushBack(1)
     x.PushBack(2)
     x.PushBack(3)

   for e := x.Front(); e != nil; e=e.Next() {
        fmt.Println(e.Value.(int))
       }
}
-------------------------------------------------




