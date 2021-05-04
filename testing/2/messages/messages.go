package messages

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func depart(name string) string {
	return fmt.Sprintf("Bye %s", name)
}
