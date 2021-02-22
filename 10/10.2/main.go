package main

import (
	"fmt"
	"time"
)
func main(){
	a:=time.Now()
	fmt.Println("Script Started at ",a)
	time.Sleep(2*time.Second)
	b:=time.Now()
	total:=b.Sub(a)
	fmt.Println("Total Duration",total.Hours())
	fmt.Println("Total Duration in seconds",total.Seconds())



}