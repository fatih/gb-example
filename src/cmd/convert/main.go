package main

import (
	"fmt"
	"os"

	"gb-example/snakecase"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "usage: convert [name]")
		os.Exit(1)
	}

	name := os.Args[1]

	// internal package
	s := snakecase.Convert(name)
	fmt.Println(s)
}
