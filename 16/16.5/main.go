package main
import "fmt"
func hello(c chan int){
	for i:=0;i<10;i++{
		c<-i
	}
	close(c)
}
func main(){
a:=make(chan int)
go hello (a)
for {
	val,status:=<-a
	if !status{
		break 
	}
	fmt.Println(val)
}
}