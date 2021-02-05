// pwd complex checker

package main

import (
	"fmt"
	"unicode"
)

func pwd(pw string) bool {
	pwR := []rune(pw) // converting string to rune
	if len(pwR) < 8{
		return false
	}
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _,v := range pwR{
		if unicode.IsUpper(v) {
			hasUpper = true

  }
  if unicode.IsLower(v) {

    hasLower = true

  }
  if unicode.IsNumber(v) {

    hasNumber = true

  }
  if unicode.IsPunct(v) || unicode.IsSymbol(v) {

    hasSymbol = true

  }


	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}
func main(){
result:=pwd("Random123@!")
fmt.Println(result)
}