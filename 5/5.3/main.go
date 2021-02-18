package main

import "fmt"

func odd (i int) (int,string){

	switch {
	case i%2==0:
		return i,"even"
	case i%2!=0:
		return i,"odd"
	
	default:
		return 0,""

	}
	
}

func main(){
for i:=0;i<10;i++{

	key,result:=odd(i)
	fmt.Println(key,result)
}
}