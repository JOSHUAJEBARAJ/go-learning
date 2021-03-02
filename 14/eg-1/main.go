package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type todo struct {
	
	Userid int `json:"userid"`
	Id int `json:"id"`
	Title  string `json:"title"`
	Completed string `json:"completed"`
	
}
func get() []todo{

	r,err:=http.Get("https://jsonplaceholder.typicode.com/todos")
	if err !=nil{
		fmt.Println(err)
	}
	response,err:=ioutil.ReadAll(r.Body)
	//fmt.Printf("%T\n",response)
	//responsestring:=string(response)
	_=err
	//valid:=json.Valid(response)
	//fmt.Println(valid)
	var todos []todo
	err1:=json.Unmarshal(response,&todos)
	if err1 !=nil{
		fmt.Println("UnMarshall error")
		fmt.Println(err)
	}

	return todos
}

func main(){
	a:=get()
	fmt.Printf("%T",a)
	// for key,value =:range a {
	// 	fmt.Println(key.Id)
	// }
	
}