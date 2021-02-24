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

	data:=[]byte(`{
		"Name":"Joshua",
		"Id":1123,
		"Courses":[{
"Coursename":"maths",
"Code":100
		},
		{
			"Coursename":"science",
			"Code":101
					}]
	}`)
var s1 student
	valid:=json.Valid(data)
	fmt.Println(valid)
	err:=json.Unmarshal(data,&s1)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(s1)

}