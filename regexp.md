## Regexp 

Find in string - !, ",", spaces and replace it with empty string

```go
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var input = "A man, a plan, a canal, Panama!" // its palendrom

	// find - "!", commas ",", spaces "\s" 
	re := regexp.MustCompile("[!,\\s]")
	
	
	replaced := re.ReplaceAllString(input, "")
	fmt.Println(replaced)

	fmt.Println(checkPhalendrom(replaced))
} 

func checkPhalendrom(value string) bool {
	var valLen = len(value)
	var result = false

	for i := 0; i < valLen; i++ {
		if strings.ToLower(string(value[i])) == strings.ToLower(string(value[(valLen-1)-i])) {
			result = true
		} else {
			result = false
		}
	}

	return result
}

```
