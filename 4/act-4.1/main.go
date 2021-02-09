package main

import "fmt"

func fillArray(arr [10] int) [10] int {
	for i:=0;i<10;i++{
		arr[i]=i+1
	}
	return arr
} 
func main(){
	var arr [10] int
	arr=fillArray(arr)
	fmt.Println(arr)
}