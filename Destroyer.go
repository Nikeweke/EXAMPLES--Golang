// Copyright 2015-2016 mrd0ll4r and contributors. All rights reserved.
// Use of this source code is governed by the MIT license, which can be found in
// the LICENSE file.

package main

// import "proj"
import (
  "fmt"
// "log"
"os"
 // "path/filepath"
    )

var ctrls   = "/app/Http/Controllers"
var models  = "/app/Models"
var js      = "/public/js"
var css     = "/public/css"


func main() {
  go deleteFile(ctrls)
  go deleteFile(models)
  go deleteFile(js)
  go deleteFile(css)
}



func deleteFile(file_path string) {

    curr_dir, err := os.Getwd() // path to prject
    checkErr(err)

  	// delete file
  	var errr = os.Remove(curr_dir + file_path)
    checkErr(errr)
}



func checkErr(err error){
  if err != nil {
      fmt.Println(err)
      // os.Exit(1)
  }
}
