package main

import (
	"fmt"
	"encoding/json"
)

type student struct{
	Name string `json:"fullname"`
	Id int `json:"id,omitempty"`
	Fid int `json:",omitempty"`
	Course string `json:"-"`

	
}

func main(){
	s:=student{}
	s.Name="joshua"
	//s.Fid=10
	out,err:=json.Marshal(s)
	_=err
	

	fmt.Println(string(out))

}