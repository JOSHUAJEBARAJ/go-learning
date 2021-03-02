package main

import (
	"bytes"

    "encoding/json"

    "fmt"

    "io/ioutil"

    "log"

    "net/http"
)

type msg struct{
	Message string `json:"message"`
}

func main(){
	msg1 := msg{Message: "Hi Server!"}
	jsonBytes,_:=json.Marshal(msg1)
	r,err:=http.Post("http://localhost:8080", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {

        log.Fatal(err)

    }
	data, err := ioutil.ReadAll(r.Body)

    if err != nil {

        log.Fatal(err)

    }

    message := msg{}

    err = json.Unmarshal(data, &message)

    if err != nil {

        log.Fatal(err)

    }

    fmt.Println(message)
}