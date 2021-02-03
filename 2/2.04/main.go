package main

import (
	"fmt"
	"errors"
)

func validate(a int) error {

	if(a<1){
		return errors.New("input can't be a negative number")

	}else if (a<100) {
		return errors.New("input can't be a big number")
	}else{
		return nil
	}
	
}
func main()  {
	a:=-1
	if err:=validate(a);err !=nil{
fmt.Println(err)
	}else{
		fmt.Println("Valid Input")
	}

}