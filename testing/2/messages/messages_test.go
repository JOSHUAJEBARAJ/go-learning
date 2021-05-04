package messages

import (
	"testing"
)

func TestGreet(t *testing.T) {
	got := Greet("Joshua")
	expected := "Hello Joshua"
	if got != expected {
		t.Errorf("Did not get the exptected value Expected :%v Got %v", expected, got)
	}

}
func TestDepart(t *testing.T) {
	got := depart("Joshua")
	expected := "Hello Joshua"
	if got != expected {
		t.Errorf("Did not get the exptected value Expected :%v Got %v", expected, got)
	}

}
