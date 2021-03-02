package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
)

func get() string{
	

	r,err:=http.Get("https://www.google.com/")
	if err !=nil{
		fmt.Println(err)
	}
	defer r.Body.Close()
	response,err:=ioutil.ReadAll(r.Body)
	_=err
	return string(response)
}

func main(){
	a:=get()
	fmt.Println(a)
}