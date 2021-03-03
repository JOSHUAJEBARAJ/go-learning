package main
import "fmt"
import "time"
func wish(s string){
	fmt.Println(s)
}

func main(){
	go wish("Hello")

	fmt.Println("Main")
time.Sleep(2*time.Second)
}