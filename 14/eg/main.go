package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	
)

func main() {
	r,err:=http.Get("https://jsonplaceholder.typicode.com/todos")
	if err!=nil{
	fmt.Println(err)
	}
	resaponsedata,err1:=ioutil.ReadAll(r.Body)
	_=err1
	fmt.Println(string(resaponsedata))
}
