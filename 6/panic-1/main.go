package main

import(
	"fmt"
	"errors"
)
var (
	lowage=errors.New("Age Should not less then 18")
	highage=errors.New("Age should not more then 35")
)

func message(a string){

	if (a=="Hello"){
		panic(errors.New("Panic"))
	}

}
func  main(){

	m:="Hello"
	message(m)
	fmt.Println("Print after panic")
	

}