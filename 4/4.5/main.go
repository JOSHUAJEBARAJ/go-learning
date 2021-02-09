package main

import "fmt"

func message() string{

    arr := [...]string{

"joshua",
"jebaraj",
    }
	arr[1]="joshua"
	arr[0]="jebaraj"
    return fmt.Sprintln(arr[0],arr[1])
}

func main(){
	fmt.Println(message())

}