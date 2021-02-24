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
	_=err
	data1:= v1.(map[string]interface{})
	for k, v := range data1 {

		switch value := v.(type) {
	
		case string:
	
		  fmt.Println("(string):", k, value)
	
		case float64:
	
		  fmt.Println("(float64):", k, value)
	
		case bool:
	
		  fmt.Println("(bool):", k, value)
	
		case []interface{}:
	
		  fmt.Println("(slice):", k)
	
		  for i, j := range value {
	
			fmt.Println(" ", i, j)
	
		  }
	
		default:
	
		  fmt.Println( "(unknown):",k, value)
	
		  }
	
	  }

}