package main

import (
	"fmt"
	"encoding/json"
)

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
var v1 interface{}
	valid:=json.Valid(data)
	fmt.Println(valid)
	err:=json.Unmarshal(data,&v1)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(v1)

}