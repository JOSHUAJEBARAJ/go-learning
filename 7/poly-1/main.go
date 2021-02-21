package main

import "fmt"
type Person struct{
	name string
}

type Animal struct{
	name string
}
type Speaker interface{
	speak() 
}
func (a Animal) speak(){
	fmt.Println("Hello Animal ",a.name)
}
func (p Person) speak(){
	fmt.Println("Hello",p.name)
}

func sayPerson(s Speaker){
	s.speak()
}
func sayAnimal(s Speaker){
	s.speak()
}
func main(){
	p:=Person{"Joshua"}
	sayPerson(p)
	a:=Animal{"Dog"}
	sayAnimal(a)

}