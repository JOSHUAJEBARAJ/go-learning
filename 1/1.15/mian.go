package main
import "fmt"

func pass_by_value(count int){

	count=count+5
	
}

func pass_by_ref(count *int){

	*count=*count+5 // while using prefix with *
	
}

func main()  {
	a:=5
	pass_by_value(a)
	fmt.Println(a)
	pass_by_ref(&a)
	fmt.Println(a)


}