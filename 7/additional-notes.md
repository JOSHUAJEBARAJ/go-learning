## Additional Notes

Interface is the type and also the value
Empty interface has no methods so all type implicity use the interface

```
package main

import (
	"fmt"
)
func empty(a interface{}){
fmt.Println(a)
}
func main() {
a:=10
empty(a)
empty("joshua")
	
}

```

https://flaviocopes.com/go-empty-interface/ 