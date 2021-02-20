// recover example

package main

import(
	"fmt"
	"errors"
)
func a (){
	
	b("Hello")
	fmt.Println("I am from A")
}

func b (m string){
	defer func() {

		if r:=recover();r!=nil{
			fmt.Println("Error",r)
		}}()
		if (m=="Hello"){
			panic(errors.New("Panic"))
		}
	}

func main(){
a()
fmt.Println("I am from main")
}