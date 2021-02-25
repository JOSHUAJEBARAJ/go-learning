package main

import(
	"fmt"
	"os"
)

func main(){
	f ,err:=os.Create("test")
	if err !=nil{
		fmt.Println(err)
	}
	f.Write([]byte("I am from write\n"))
	f.WriteString("I am from write string\n")
	defer f.Close()
}