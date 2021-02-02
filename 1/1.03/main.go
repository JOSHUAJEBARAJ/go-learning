// Exercise 1.01: Using Variables, Packages, and Functions to Print Stars

package main

import ( // remember the bracket
	"fmt"
	"time"
	
)
var (
	Debug   bool   = false

  LogLevel  string  = "info"

  startUpTime time.Time = time.Now()
)
func main(){
	fmt.Println(Debug, LogLevel, startUpTime)


}