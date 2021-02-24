package main

import (
	"fmt"
	"encoding/json")

	type student struct {
		N string `json:"name"`
	}

func main(){

	var s student

	d:=[]byte(`
	{
		"name":"joshua"
	}
`)

err:=json.Unmarshal(d,&s)
if err !=nil{
	fmt.Println(err)
}
fmt.Println(s.N) // different name then the json field

}