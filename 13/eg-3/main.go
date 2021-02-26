package main

import "fmt"

import "database/sql"

import _ "github.com/lib/pq"

func main(){
	var id int =1

var name string
db, err := sql.Open("postgres", "user=user password=pass host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
if err !=nil{
	fmt.Println(err)
}else{
	fmt.Println("Connection made")
}
connectivity := db.Ping()

if connectivity != nil{

	panic(err)
  
  }else{
  
	fmt.Println("Good to go!")
  
  }

  qryrow, err := db.Prepare("SELECT name FROM student WHERE id=$1")
  

	err = qryrow.QueryRow(id).Scan(&name)

if err != nil {

  panic(err)

}

fmt.Println("The name column value is",name,"of the row with id=",id)

qryrow.Close()
 defer db.Close()
}