package main

import "fmt"

func fullname () (string,string){
	return "joshua","jebaraj"
  }
func main()  {
	a,b :=fullname() // return two value
	fmt.Println(a,b)
}