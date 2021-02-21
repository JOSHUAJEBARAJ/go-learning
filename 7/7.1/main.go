package main

import "fmt"

type Speaker interface{
	Speak () string
}
type person struct{
	name string
	age int
}
func main(){

	p:=person{"joshua",19}
	fmt.Println(p.Speak()) // interface in the current package
	fmt.Println(p)


}

func (p person) Speak() string{
	return "Hello "+p.name
}

func (p person) String() string{
	return p.name
}