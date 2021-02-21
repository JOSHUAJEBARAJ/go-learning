package main

import "fmt"
type Dog struct{}
type Cat struct{}
type Speaker interface{

	speak()
	
}
func (a Dog)speak(){
	fmt.Println("Bow Bow")
}
func (b Cat)speak(){
	fmt.Println("meow Meow")
}

func main(){
a:=Dog{}
b:=Cat{}

makeSound(a,b)

}

func makeSound(s...Speaker){
	for _,key:=range s{
key.speak()
	}
}