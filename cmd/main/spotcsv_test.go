package main

import "testing"

func ExampleMain() {
	goMain([]string{"spocsv"})
	// Output:
	// Hello, World!
}

func TestHello(t *testing.T) {
	got := hello()
	wont := "Hello, World!"
	if got != wont {
		t.Fatalf("hello() wont %s, but got %s", wont, got)
	}
}
