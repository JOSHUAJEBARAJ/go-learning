// Modifying the Contents of an Array in a 
package main

import "fmt"
func fillArray(arr [10] int)  [10]int{
	for i:=0;i<10;i++{
		arr[i]=i + 1
	}

	return arr
}

func opArray(arr [10] int)  [10]int{
	for i:=0;i<10;i++{
		arr[i]=arr[i] * arr[i]
	}

	return arr
}
func main()  {
	
	var arr [10]int
	arr = fillArray(arr)
	arr = opArray(arr)

  fmt.Println(arr)

}