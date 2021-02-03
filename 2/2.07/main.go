package main

import(
	"fmt"
	"time"
)
func main()  {
	a:=time.Monday 
	switch { 
	case a==time.Monday:// case
		fmt.Println("Born on week starting")
	case a==time.Tuesday:
		fmt.Println("Born on week Ending")
	

	}
}