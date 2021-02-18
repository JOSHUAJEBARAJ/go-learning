package main

import "fmt"
func itemSold(){
	sales:=map[string] int {
		"Joshua":10,
		"Jebaraj":20,
	}

	for key,value:=range(sales){
		fmt.Printf("%s sold %d items \n",key,value)
	}
}
func main(){
itemSold()
}