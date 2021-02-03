package main

// break and continue
import (
	"fmt"
	"math/rand"
)

func main()  {
for{
	r:=rand.Intn(8)
	fmt.Println(r)
	if(r%3==0){
		fmt.Println("Continue")
		continue
	}
	if(r%2==0){
		fmt.Println("break")
		break
	}
}	
}