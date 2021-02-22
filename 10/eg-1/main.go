package main

import (
	"fmt"
	"time"
)
func main(){

	start:=time.Now()
	fmt.Println("Program started at ",start)
	time.Sleep(2*time.Second)
	end:=time.Now()
	fmt.Println("Program ended at ",end)




}