package main

import(
	"fmt"
	"errors"
)
var (
	lowage=errors.New("Age Should not less then 18")
	highage=errors.New("Age should not more then 35")
)

func validate(a int) (string,error){
	if (a<18){
		return "",lowage
	}else if (a>35){
		return "",highage
	}else{
		fmt.Println("I am called")
		return "valid",nil
	}

}
func  main(){
	age1,age2,age3:=16,21,40
	_=age3
	result,err:=validate(age1)
	if err !=nil{
		fmt.Println(err)
		fmt.Println(result)
	}else{
		fmt.Println(result)
	}
// we have to check 

	result1,err1:=validate(age2)
	if err !=nil{
		fmt.Println(err1)
	}else{
		fmt.Println(result1)
	}

}