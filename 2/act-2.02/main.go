package main

import "fmt"

func main()  {
	words := map[string]int{

		"Gonna": 3,
	  
		"You": 3,
	  
		"Give": 2,
	  
		"Never": 1,
	  
		"Up":  4,
	  
		}
		big_value:=0
		big_word:=""
		
		for key,value:=range words{

			if (big_value<value){
				big_value=value
				big_word=key
				

			}
		}
		fmt.Println("Most popular word:"+big_word)
		fmt.Printf("With a count of :%d",big_value)
}