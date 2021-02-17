package main

import(
	"fmt"
	"os"
)
func getPassedArgs() [] string{


	var args []string 
		for i :=1;i<len(os.Args);i++{
			args = append(args, os.Args[i]) // first parameter the slice and the second parameter the value


		
		
	}
	return args 
}

func getLocals(getPassedArgs()) {

	var longest string

  for i := 0; i < len(args); i++ {

    if len(args[i]) > len(longest) {

      longest = args[i]

  }

  }

  return longest



}
func main(){

	if longest := findLongest(getPassedArgs(3)); len(longest) > 0 { // one liner if

		fmt.Println("The longest word passed was:", longest)
	
	  } else {
	
		fmt.Println("There was an error")
	
		os.Exit(1)
	
	  }
	
	}
