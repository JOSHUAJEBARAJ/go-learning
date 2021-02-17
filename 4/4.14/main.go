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
	users := getUser()
    id:="01"
	user, exists := users[id]
	if exists{

		fmt.Println(user)
	}else
	{
		fmt.Printf("User id %s not found",id)
	}


}