package main

import "fmt"

import "database/sql"

import _ "github.com/lib/pq"

func main(){
	var id int

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

  rows, err := db.Query("SELECT * FROM student")

  for rows.Next() {

	err := rows.Scan(&id, &name)
  
	if err != nil {
  
	  panic(err)
  
	}
  
	fmt.Println(id, name)
  
  }

rows.Close()
 defer db.Close()
}