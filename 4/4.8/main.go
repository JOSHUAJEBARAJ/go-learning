package main

import(
	"fmt"
	"os"
)
func getPassedArgs(min int) [] string{

	if  len(os.Args)< min  {

		fmt.Printf("Minimum %d arguments are needed",min)
		os.Exit(1) // return with error

	}
	var args []string 
		for i :=1;i<len(os.Args);i++{
			args = append(args, os.Args[i]) // first parameter the slice and the second parameter the value


		
		
	}
	return args 
}

func findLongest(args []string ) string {

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
