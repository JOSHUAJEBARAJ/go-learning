package main
import "fmt"
func main()  {
	var at float32=0.075
	var bt float32=0.2
	var mt float32 = 0.015
	var a float32=0.99
	var m float32=2.75
	var b float32 = 0.87
	var total float32
	total=(at*a)+(bt*b)+(mt*m)
	fmt.Println(total)
	
	
}