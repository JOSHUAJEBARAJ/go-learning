package main

import (
	"fmt"
	"time")

	func whattime() string{
		return time.Now().Format(time.ANSIC)
	}
func main()  {
	fmt.Println(whattime())
	
}