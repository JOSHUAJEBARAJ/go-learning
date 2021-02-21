package main

import "fmt"
type Person struct{
	name string
}
type Speaker interface{
	speak() 
}

func (p Person) speak(){
	fmt.Println("Hello",p.name)
}

func say(s Speaker){
	s.speak()
}
func main(){
	p:=Person{"Joshua"}
	say(p)

}