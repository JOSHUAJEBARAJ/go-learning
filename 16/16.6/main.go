package main
import "fmt"
func main(){
	a:=make(chan int,2) // creating buffer
	a<-10
	a<-20
	fmt.Println(<-a)
	fmt.Println(len(a))
	fmt.Println(<-a)

}