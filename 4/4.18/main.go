package main

import "fmt"
type rectangle struct{
	length  int 
	breadth int 
}
func main(){
	rectangle1:= struct{
		length int 
		breadth int 
	}{
		10,
		10,
	}
	rectangle2:= struct{
		length int 
		breadth int 
	}{}
	rectangle2.length=10
	rectangle2.breadth=10
	fmt.Println(rectangle1==rectangle2)

}