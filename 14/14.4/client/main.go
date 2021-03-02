package main

import (

    "fmt"

    "io/ioutil"

    "log"

    "net/http"

    "time"

)

func main(){
	client := http.Client{Timeout: 11 * time.Second}
	req, err := http.NewRequest("POST", "http://localhost:8080", nil)

    if err != nil {

        log.Fatal(err)

    }
	req.Header.Set("Authorization", "superSecretToken")
	resp, err := client.Do(req)

    if err != nil {

        log.Fatal(err)

    }

	data, err := ioutil.ReadAll(resp.Body)

    if err != nil {

        log.Fatal(err)

    }
	fmt.Println(string(data))
}