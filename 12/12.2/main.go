package main

import (
	"fmt"
	"os"
	"io/ioutil"
)
func main(){

	// reading the file
	content,err:=ioutil.ReadFile("test")
	if err !=nil{
		fmt.Println(err)
	}
	create_error:=ioutil.WriteFile("backup", content, 0644)
	if create_error !=nil{
		fmt.Println(err)
	}
	f,_:=os.OpenFile("test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.Write([]byte("New"))
	defer f.Close()


}