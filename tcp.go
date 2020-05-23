package main

import (
  "log"
  "fmt"
  "net"
  "time"
)

// var TCP_ADDRESS = "192.168.0.130:7110"
var TCP_ADDRESS = "92.60.180.73:7110"
// var TCP_ADDRESS = "92.60.180.73:8234"

func main() {

  fmt.Println("TCP =", TCP_ADDRESS)
  conn, err := net.Dial("tcp", TCP_ADDRESS)
  // Подключаемся к сокету
  if err != nil {
    log.Fatal(err)
    return
  }


  buf := make([]byte, 10)

  for {
    // set SetReadDeadline
    err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
    if err != nil {
      log.Println("SetReadDeadline failed:", err)
      // do something else, for example create new conn
      return
    }

    writeData := []byte("PZ")
    fmt.Println("Data sent -", writeData)

    n, err := conn.Write(writeData)
    if err != nil {
      log.Fatal(err)
      return 
    }
    fmt.Printf("Bytes sent - %d \n\n", n)
  
    conn.Read(buf)
    log.Println("String buf -", string(buf))
    log.Println("Raw buf -", buf, "\n")
  }
}

