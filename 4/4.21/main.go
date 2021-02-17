package main

import (
	"fmt"
	"errors"
)
func doubler(a interface {})(string,error){
	if i, ok := a.(int); ok{
		return fmt.Sprint(i * 2), nil


	}
	if i, ok := a.(string); ok{
		return fmt.Sprint(i+i), nil


	}
	return "", errors.New("unsupported type passed")
}
func main(){
	res,_:=doubler(5)
	fmt.Print(res)
	_,err:=doubler(true)
	fmt.Print(err)

}