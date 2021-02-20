package main

import(
	"fmt"
	
)


func main(){
a:=10
b:="Joshua"
_=b
fmt.Printf("a= %v and its type %T \n",a,a)
fmt.Printf("a= %v and its type %[1]T \n ",a)

}
