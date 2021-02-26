package main


import "fmt"

import "database/sql"

import _ "github.com/lib/pq"
var id int =1
func main(){


db, err := sql.Open("postgres", "user=user password=pass host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
if err != nil {

	panic(err)
	
	}else{
	
	  fmt.Println("The connection to the DB was successfully initialized!")
	
	}

	test1, _ := db.Exec("TRUNCATE TABLE test")
test, _ := db.Exec("DROP TABLE test")

test=test1
_=test

db.Close()

}
