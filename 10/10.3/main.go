package main

import (
	"fmt"
	"time"
)

func findTime(a string) string{
	cur:=time.Now()
zone,err:=time.LoadLocation(a)
if err !=nil{
	fmt.Println(err)
	fmt.Println("Wrong Zone")
}
rt:=cur.In(zone)
return rt.Format(time.ANSIC)
}
func main(){

	fmt.Println("Time in America",findTime("America/Los_Angeles"))

	



}