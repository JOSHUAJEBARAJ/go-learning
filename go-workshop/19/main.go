package main

import "fmt"
type  student struct{
	name string
	age int
}

// similar to declaration
func (s *student) ismajor(){

	if s.age>3{
		fmt.Println("allow")
	}
}
func main(){
s:=student{"Joshua",10}
s.ismajor()
}