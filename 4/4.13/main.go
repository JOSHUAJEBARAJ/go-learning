package main

import "fmt"

func getUser() map[string] string{

	users := map[string]string{
		"02": "Joshua",
		"01": "Ezhil",

	}

	users["03"]="Surya" // adding the new element
	return users
}

func main(){

fmt.Println("Users",getUser())
}