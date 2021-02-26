package main

import "fmt"

import "database/sql"

import _ "github.com/lib/pq"

func main(){
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

//   create:=`CREATE TABLE student(
  
// 	id integer NOT NULL,
  
// 	name text NOT NULL
  
//   );
  
//   `
//   _,err = db.Exec(create)

//   if err != nil {
  
// 	panic(err)
  
//   } else{
  
// 	fmt.Println("The table called Number was successfully created!")
  
//   }
s:=[]string{"Joshua","Jebaraj"}
  insert, insertErr := db.Prepare("INSERT INTO student VALUES($1,$2)")
  if insertErr != nil{

	panic(insertErr)
  
  }
  for key,value:=range s{
	insert.Exec(key,value)
  }
  insert.Close()
 defer db.Close()
}