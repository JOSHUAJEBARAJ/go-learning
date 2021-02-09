package main

import "fmt"
func main()  {
	var arr1 [5]int

	arr2 := [5]int{0}
  
	arr3 := [...]int{0, 0, 0, 0, 0}
	
	comp1:=arr1 == arr2
	comp2:=arr2 == arr3
	fmt.Println(comp1,comp2)
}