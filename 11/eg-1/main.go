package main

import (
	"fmt"
	"encoding/json"
)
type student struct {
	Name string // must be caps
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
fmt.Println(s.Name)
}