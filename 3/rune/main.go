package main

import "fmt"

func main() {
	username := "Sir_King_Über"
	for i := 0; i < len(username); i++ {

		fmt.Print(username[i], " ")
	// 83 105 114 95 75 105 110 103 95 195 156 98 101 114 There are only 13 char but we have 14 byte                                                          
	  }
	  for i := 0; i < len(username); i++ {

		fmt.Print(string(username[i]), " ")
	// 83 105 114 95 75 105 110 103 95 195 156 98 101 114 There are only 13 char but we have 14 byte                                                          
	  }
	
	  runes := []rune(username)

	  for i := 0; i < len(runes); i++ {
	    fmt.Print((runes[i]), " ")
		//fmt.Print(string(runes[i]), " ")
	
	  }}
	                                                 