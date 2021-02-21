package main

import "fmt"
func main(){

	a:=[]int{1,2,3,4}


	b:=a[2:3]
	fmt.Printf("%d %d %v \n",cap(b),len(b),b[0])
	make1:=make([]int,5)
	//fmt.Println(b[1])
	//make2:=make([]int,0,5)
	fmt.Printf("%d %d \n",cap(make1),len(make1))
	

}