package main

import(
	"fmt"
	"time"
)
func main()  {
	a:=time.Monday 
	switch a{ // using variable as expression
	case time.Monday:
		fmt.Println("Born on monday")
	case time.Tuesday:
		fmt.Println("Born on Tuesday")
	

	}
}