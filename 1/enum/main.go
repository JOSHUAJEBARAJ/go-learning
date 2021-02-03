package main


import "fmt"

const (
	a=iota // starts with 0
	b
	c
)

//  to start with 1 use a=iota+1
func main()  {
fmt.Println(a,c)	
}