package main

import(
	"fmt"
	"errors"
)
func divide(a,b int) (int ,error){
	if b==0{
		return 0,errors.New("Cannot divided by zero")
	}else{
		return b/a,nil
	}
}

func main(){
result,err:=divide(2,0)
if err !=nil{
	fmt.Println(err)
}else{
	fmt.Println(result)
}
}
