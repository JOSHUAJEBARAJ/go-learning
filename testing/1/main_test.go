package main_test

import "testing"

func TestAddition(t *testing.T) {

	got := 2 + 2
	expected := 4
	if got != expected {

		t.Errorf("didnt get the expected result %v", got)
	}
}

func TestSubAddition(t *testing.T) {

	got := 2 + 3
	expected := 4
	if got != expected {

		t.Errorf("didnt get the expected result %v", expected)
	}
}
