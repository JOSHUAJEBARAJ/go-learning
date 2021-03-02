package main

import(
	"log"
)
func greet(c chan string ){
	c <-"hello"

}

func main(){
	m:=make(chan string)
	go greet (m)
	log.Println(<-m)
}