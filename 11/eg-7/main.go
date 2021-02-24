package main

import (
	"fmt"
	"encoding/json"
)

type student struct{
	Name string `json:"fullname"`
	Id int `json:"id,omitempty"`
	
}

func main(){
	s:=student{}
	s.Name="joshua"
	s.Id=1123
	out,err:=json.MarshalIndent(s,""," ")
	_=err
	

	fmt.Println(string(out))

}