package main

import "fmt"
import "time"
func sendonly(c chan<- int){
	fmt.Println("Iam called")
c<-10
}

func main(){
	a:=make(chan<- int)
	go sendonly(a)
	b:=make(chan int)
	go sendonly(b)
	time.Sleep(10*time.Second)
}