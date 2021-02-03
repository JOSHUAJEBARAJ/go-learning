package main

// looping over array
import "fmt"

func main(){
	names := []string{"Jim", "Jane", "Joe", "June"}
	for i:=0;i<len(names);i++{ // no curves	
		fmt.Println(names[i])
	}
}