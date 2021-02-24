package main

import (
	"fmt"
	"encoding/json"
)
type student struct{
	Name string
	Id int 
	Courses []course
}
type course struct{
	Coursename string 
	Code int 
}
func main(){

	
var s1 student
	s1.Name="Joshua"
	s1.Id=100
	

}