package main


import (
	"net"
	"fmt"
	"encoding/base64"
	"encoding/xml"
)

var HOST = "*********"
var PORT = ""*********""
var USER = ""*********""
var PWD = ""*********""

var TCP_HOST = HOST + ":" + PORT
var FlagForOnce = 0

type EventNotificationAlert struct {
	BlockName   xml.Name `xml:"EventNotificationAlert"`
	IpAddress   string   `xml:"ipAddress"`
	PortNo      string   `xml:"portNo"`
	Protocol    string   `xml:"protocol"`
	MacAddress  string   `xml:"macAddress"`
	ChannelID   string   `xml:"channelID"`
	DateTime   string    `xml:"dateTime"`
	ActivePostCount   string   `xml:"activePostCount"`
	EventType   string   `xml:"eventType"`
	EventState   string   `xml:"eventState"`
	EventDescription   string   `xml:"eventDescription"`
}


func main() {

  var basicAuthToken = "Basic " + basicAuth(USER, PWD)

	var request = "GET /some/route HTTP/1.1\r\n" +
	"Host: " + TCP_HOST + "\r\n" +
	"Content-Type: application/xml \r\n" +
	"Authorization: " + basicAuthToken + "\r\n" + 
	"Accept: multipart/x-mixed-replace\r\n\r\n";

  fmt.Println(request)

	// Подключаемся к сокету
	conn, err := net.Dial("tcp", TCP_HOST)
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	// Sending in socket HTTP data - only 1 time, 
	// that will establish connection for receiving current alarms events
	fmt.Fprintf(conn, request)

  for { 
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}

		response := string(buf[:len])

    // parse response to struct
		data := &EventNotificationAlert{}
		xml.Unmarshal([]byte(response), &data)
		
		fmt.Println(data.EventType)

		if data.EventType == "vehicledetection" {
			fmt.Println(response)
		}

		// fmt.Println("===========================>")
		// fmt.Println("Stuff", s)
		// fmt.Println("len", binary.Size(buf))

  }
}


func basicAuth(username, password string) string {
  auth := username + ":" + password
   return base64.StdEncoding.EncodeToString([]byte(auth))
}


