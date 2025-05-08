package main

import (
	"testing"

	"github.com/tamada/spotcsv/cmd/main/commands"
)

func Example() {
	goMain([]string{"spocsv"})
	// Output:
	// Usage:
	//   spotcsv [flags]
	//   spotcsv [command]
	//
	// Available Commands:
	//   completion  Generate the autocompletion script for the specified shell
	//   except      Except specified indexes items from the CSV
	//   head        Show specified lines from head in the CSV
	//   help        Help about any command
	//   index       Select specified indexes items from the CSV
	//   lines       Show specified lines in the CSV
	//   tail        Show specified lines from tail in the CSV
	//
	// Flags:
	//   -h, --help   help for spotcsv
	//
	// Use "spotcsv [command] --help" for more information about a command.
}

func TestNoArgumentsMain(t *testing.T) {
	err := commands.Execute([]string{"spocsv"})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "no command specified" {
		t.Fatalf("expected 'no command specified', got %s", err.Error())
	}
}

// func TestHello(t *testing.T) {
// 	got := hello()
// 	wont := "Hello, World!"
// 	if got != wont {
// 		t.Fatalf("hello() wont %s, but got %s", wont, got)
// 	}
// }
