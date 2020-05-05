package main

import (
	"fmt"
	"os"
)

var version string = "dev"

func main() {
	err := run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(args []string) error {
	fmt.Printf("I am running the %s version!\n", version)
	return nil
}
