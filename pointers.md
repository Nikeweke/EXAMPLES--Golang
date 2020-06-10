### Pointer

A pointer is a variable whose value is the address of another variable, i.e., direct address of the memory location.

```go
package main

import (
	"fmt"
)

func main() {
	var a int = 20   /* actual variable declaration */
   var ip *int      /* pointer variable declaration */

   ip = &a  /* store address of a in pointer variable*/

   fmt.Printf("Address of a variable: %x\n", &a  )

   /* address stored in pointer variable */
   fmt.Printf("Address stored in ip variable: %x\n", ip )

   /* access the value using the pointer */
   fmt.Printf("Value of *ip variable: %d\n", *ip )
}
```
