package main


import (
	"fmt"
	"os"
)

func main(){
	file,err:=os.Stat("not")
	if err !=nil{
		if os.IsNotExist((err)) {

			fmt.Println("not: File does not exist!")
		
			fmt.Println(file)
		
		  }
	}
}