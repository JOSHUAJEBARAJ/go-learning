package main

import(
	"fmt"
	"errors"
)


func message(a string){
	m1:=func(){
		fmt.Println("Defer inside the function")
	}
	defer m1()

	if (a=="Hello"){
		panic(errors.New("Panic"))
	}

}
func  main(){

	m:=func(){
		fmt.Println("Defer outstide the function")
	}
	defer m()
	me:="Hello"
	message(me)
	

}