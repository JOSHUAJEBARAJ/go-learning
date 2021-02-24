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
	out,err:=json.Marshal(s)
	_=err
	

	fmt.Println(string(out))

}