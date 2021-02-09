package main

import "fmt"
func message() string{

    arr := [...]string{

"joshua",
"jebaraj",
    }
    return fmt.Sprintln(arr[0],arr[1])
}

func main()  {
	fmt.Print(message())
}