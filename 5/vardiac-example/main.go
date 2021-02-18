package main

import "fmt"
func print(i ... int){
fmt.Println(i)
fmt.Printf("%T",i)
}
func main(){

	print(1,2)
	a:=[]int{1,3,4} // cannot use a (type []int) as type int in argument to print
	print(a...)

}