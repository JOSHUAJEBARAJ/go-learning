package main

import "fmt"

type student struct{
	name string
	no int
}

func getUsers() []student{
	s1:=student{
		name:"joshua", // use semicolon
		no:1123,
	}
	s2:=student{
		"anand",
		1040,
	}
	var s3 student
	s3.name="Ezhil"
	s3.no=102

	return []student{s1,s2,s3}

}
func main(){
students:=getUsers()
for i:=0 ; i < len(students); i++ {
	fmt.Println(students[i])
}
}