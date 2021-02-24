package main

import (
	"fmt"
	"encoding/json"
)

type student struct{
	Name string `json:"fullname"`
}

func main(){
	s:=student{"Joshua"}
	out,err:=json.Marshal(s)
	_=err
	//fmt.Println((out)) // print the byte

	fmt.Println(string(out))

}