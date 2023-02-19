### Read by chunks

###### Request body

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)


// Получение даты из файла
func ReadRequestBodyByChunks(url string) (string, error) {
	// set chunkSize in bytes
	chunkSize := int64(120)
	maxBytesToRead := chunkSize // change this one if want to increase size of read content
	nBytes, nChunks := int64(0), int64(0)
	buf := make([]byte, 0, chunkSize)
	result := ""

	// making http request
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("ERROR:" + err.Error())
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR:CLIENT:" + err.Error())
		return "", err
	}
	defer res.Body.Close()

	// creating reader based on request body
	r := bufio.NewReader(res.Body)

	// reading by chunks
	for {
		// if read text exceed maximum then break for loop
		if nBytes >= maxBytesToRead {
			break
		}

		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]

		if n == 0 {
			log.Println(err)
			return "", err
		}
		if n == 0 && err == nil {
			continue
		}
		if n == 0 && err == io.EOF {
			break
		}

		// adding read strings 
		result += strings.Trim(string(buf), " ")

		nChunks++
		nBytes += int64(len(buf))

		// process buf
		if err != nil && err != io.EOF {
			log.Println(err)
			return "", err
		}
	}

	
	return result, nil
}
```


###### File

```go
func ReadFileByChunks(url string) (string, error) {
	// open file 
	fileName := "data.json"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", fileName, err.Error())
	}
	// creating reader based on request body
	r := bufio.NewReader(file)

	// set chunkSize in bytes
	chunkSize := int64(120)
	maxBytesToRead := chunkSize // change this one if want to increase size of read content
	nBytes, nChunks := int64(0), int64(0)
	buf := make([]byte, 0, chunkSize)
	result := ""

	// reading by chunks
	for {
		// if read text exceed maximum then break for loop
		if nBytes >= maxBytesToRead {
			break
		}

		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]

		if n == 0 {
			log.Println(err)
			return "", err
		}
		if n == 0 && err == nil {
			continue
		}
		if n == 0 && err == io.EOF {
			break
		}

		// adding read strings 
		result += strings.Trim(string(buf), " ")

		nChunks++
		nBytes += int64(len(buf))

		// process buf
		if err != nil && err != io.EOF {
			log.Println(err)
			return "", err
		}
	}

	return result, nil
}
```
