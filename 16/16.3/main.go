package main
import "fmt"
func hello(c chan int){
fmt.Println("Hello")
c<-10
}
func main(){
	
	c:=make(chan int)
	go hello(c)
	<-c  // it will wait until reads operation is done the further code will not be executed
	fmt.Println("Main")
}