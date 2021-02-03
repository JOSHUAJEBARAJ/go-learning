// getting value from the pointer

package main

import "fmt"
func main()  {
	var count *int // default value is 0
	if count != nil{
fmt.Printf("%#v",*count)
	}
	count2:=new(int) // 0 is the default
	if count2 != nil{
		fmt.Printf("%#v",*count2)
			}
}