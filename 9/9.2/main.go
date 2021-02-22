// Printing Decimal, Binary, and Hex Values

package main

import "fmt"
func main(){

	for i:=1;i<255;i++{
		fmt.Printf("%10.d %15.b %10.x\n",i,i,i)
	}

}