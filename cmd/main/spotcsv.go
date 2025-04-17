package main

import (
	"fmt"
	"os"
)

func hello() string {
	return "Hello, World!"
}

func goMain(args []string) int {
	fmt.Println(hello())
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
