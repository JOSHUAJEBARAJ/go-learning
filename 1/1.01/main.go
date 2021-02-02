// Exercise 1.01: Using Variables, Packages, and Functions to Print Stars

package main

import ( // remember the bracket
	"fmt"
	"math/rand"
    "strings"
   "time"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1 // to get the random numher
	stars := strings.Repeat("*", r) // repeating
	fmt.Println(stars)


}