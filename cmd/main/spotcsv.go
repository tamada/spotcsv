package main

import (
	"os"

	"github.com/tamada/spotcsv/cmd/main/commands"
)

func goMain(args []string) int {
	commands.Execute(args)
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
