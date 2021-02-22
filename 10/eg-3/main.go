package main

import (
	"fmt"
	"time"
)
func main(){

	t1, _ := time.Parse(time.RFC3339,"2019-09-27T22:18:11+00:00")

	t2, _ := time.Parse(time.UnixDate,"2019-09-27T22:18:11+00:00")
  
	t3, _ := time.Parse(time.ANSIC,"2019-09-27T22:18:11+00:00")
  
	fmt.Println("RFC3339:",t1)
  
	fmt.Println("UnixDate",t2)
  
	fmt.Println("ANSIC",t3)



}