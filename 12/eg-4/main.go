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
	defer f.Close()
}