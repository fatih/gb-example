package main

import (
	"fmt"

	"github.com/fatih/gb-example/util"
)

func main() {
	fmt.Println("An example gb project")

	// internal package
	s := util.SnakeCase("GbExample")
	fmt.Println(s)
}
