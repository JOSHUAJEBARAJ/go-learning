package main

import (
	"fmt"
	"time")

func main(){
	today:=time.Now().Weekday()
   switch today {
   case time.Saturday:
	fallthrough
   case time.Sunday:
	fmt.Println("Its weekend")
   case time.Friday:
	fmt.Println("Sooon")
   default:
	fmt.Println("Too long")
	   
   }
}