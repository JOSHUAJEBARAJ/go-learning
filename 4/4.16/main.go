package main

import "fmt"
type name string 
func main(){

	var name1 name ="joshua" 
	var name2 name = "joshua"
	fmt.Println(string(name2)=="joshua") // compare it with the base type
	fmt.Println(name1==name2) // comparing it with the other type

}