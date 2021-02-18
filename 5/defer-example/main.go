package main

import "fmt"

func main(){

	defer done() // 3rd
	defer func (){
		fmt.Println("Anonmouyous defer") // 2nd
	}()
	fmt.Println("Joshua") // 1 st
}

func  done()  {
	fmt.Println("Last")
}