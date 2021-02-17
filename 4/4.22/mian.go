package main

import (
	"fmt"
	"errors"
)
func doubler(a interface {})(string,error){
	switch t:=a.(type){
	case string:
		return t+t,nil
	
	default:

	return "", errors.New("unsupported type passed")
  
	}
}
func main(){
	res,_:=doubler("Jos")
	fmt.Print(res)
	_,err:=doubler(true)
	fmt.Print(err)

}