package main


import "fmt"

import "database/sql"

import _ "github.com/lib/pq"
var id int =1
func main(){

UpdateStatement :=`

UPDATE student

SET name = $1

WHERE id = $2

`
db, err := sql.Open("postgres", "user=user password=pass host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
if err != nil {

	panic(err)
	
	}else{
	
	  fmt.Println("The connection to the DB was successfully initialized!")
	
	}

UpdateResult, UpdateResultErr := db.Exec(UpdateStatement,"dinesh",id)

if UpdateResultErr != nil {

  panic(UpdateResultErr)

}

UpdatedRecords, UpdatedRecordsErr := UpdateResult.RowsAffected()

if UpdatedRecordsErr != nil {

  panic(UpdatedRecordsErr)

}

fmt.Println("Number of records updated:",UpdatedRecords)

db.Close()

}
