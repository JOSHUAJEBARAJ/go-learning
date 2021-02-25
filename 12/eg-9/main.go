package main
import ("fmt"
"io/ioutil"

  "os")
func main(){
	f,err:=os.Open("test")
	if err !=nil{
		fmt.Println(err)
	}
	content,_:=ioutil.ReadAll(f)
	fmt.Println(string(content))

}