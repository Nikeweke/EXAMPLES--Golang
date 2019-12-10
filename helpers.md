## Useful funcs

### Get random number for some array in his range of size
```go
      /*
      |--------------------------------------------------------------------------
      |   Random number of array
      |--------------------------------------------------------------------------
      */
      func GetRandomNumb(array []string) int{
        s         := rand.NewSource(time.Now().Unix())
        r         := rand.New(s) // initialize local pseudorandom generator
        return    r.Intn(len(array))
      }
      
      // using 
       wicther := [...]string{"...", "..."}    
       rand_numb := GetRandomNumb(witcher[:])
```


### CheckErr
```go
    func checkErr(err error){
      if err != nil {
          fmt.Println(err)
          // os.Exit(1)
      }
    }
```


### Get full path to project
```go
func main(){
curr_dir, err := os.Getwd()
checkErr(err)
fmt.Prinln(curr_dir)
}
```
