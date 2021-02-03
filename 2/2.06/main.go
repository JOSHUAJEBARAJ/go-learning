package main

import(
	"fmt"
	"time"
)
func main()  {
	a:=time.Monday 
	switch a{ // using variable as expression
	case time.Monday,time.Tuesday,time.Wednesday:
		fmt.Println("Born on week starting")
	case time.Thursday,time.Friday:
		fmt.Println("Born on week Ending")
	

	}
}