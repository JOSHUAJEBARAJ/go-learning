package main


import "fmt"
func main() {
	logLevel := "デバッグ"
	fmt.Printf("%T \n",logLevel)
	for index, runeVal := range logLevel {
		fmt.Println(index, string(runeVal))
	}}

