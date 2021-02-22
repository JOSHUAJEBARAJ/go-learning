package main

import (
	"fmt"
	"time"
)
func main(){

	today:=time.Now()
	tommorow:=today.Add(24*time.Hour)
	g1:=today.Before(tommorow)
	g2:=tommorow.After(today)
	t2:=today
	fmt.Println(g1)
	fmt.Println(g2)
	fmt.Println(today.Equal(t2))



}