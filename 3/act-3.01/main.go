package main
import "fmt"
func main()  {
	var at float32=0.075 // cake tax
	var bt float32=0.02 // butter tax
	var mt float32 = 0.015 //milk tax
	var a float32=0.99 // cake
	var m float32=2.75 // milk
	var b float32 = 0.87 // butter
	var total float32
	total=(at*a)+(bt*b)+(mt*m)
	fmt.Println(total)
	
	
}