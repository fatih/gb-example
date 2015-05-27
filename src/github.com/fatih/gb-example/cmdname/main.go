package main

import (
	"fmt"

	"github.com/fatih/gb-example/snakecase"
)

func main() {
	fmt.Println("An example gb project")

	// internal package
	s := snakecase.Convert("GbExample")
	fmt.Println(s)
}
