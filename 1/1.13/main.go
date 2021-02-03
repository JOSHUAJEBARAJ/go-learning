package main

import (
	"fmt"
	"time"
)
func main()  {
	var a *int // method-1
	b:=new(int) //method-2
	c:=10
	d:=&c // method-3
	t := &time.Time{} // method-4
	fmt.Printf("%#v",a)
	fmt.Printf("%#v",b)
	fmt.Printf("%#v",d)
	fmt.Printf("%#v",t)

	
}