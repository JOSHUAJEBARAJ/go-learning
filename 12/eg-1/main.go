package main

import (
	"fmt"
	"flag"
)
func main(){

	name:=flag.String("name","","Enter the username")
	flag.Parse()
	fmt.Println(*name)

}