package main

import "fmt"

type name struct{
	fname string
	lname string
}
type student struct{
	name 
	age int 
}
func main(){
var a student // creating using the var keyword
 b:=student{}
 b.fname="joshua"
 b.lname="jebaraj"
 b.age=12

 c:=student{
	 name: name{
		 fname:"joshua",
		 lname:"jebaraj",
	 },
	 age:13,
 }

 d:=student{} // extendend method of of b
 d.name.fname="joshua"
 d.name.lname="jebaraj"
 d.age=14
students :=[] student{a,b,c,d}

for i:=0; i<len(students) ;i++{
	fmt.Println(students[i])
}
}