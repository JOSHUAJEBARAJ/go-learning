package main
import "fmt"
func main(){
	var c chan int
	if c==nil{
		fmt.Println("Not intialized")
	}
	c=make(chan int )
}