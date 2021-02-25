package main

import (
	"fmt"
	"flag"
	"os"
)
func main(){

	name:=flag.String("name","","Enter the username")
	flag.Parse()
	if *name == ""{
		flag.PrintDefaults()

  os.Exit(1)

	}
	fmt.Println(*name)

}