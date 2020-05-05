package main

import (
	"fmt"
	"os"
)

func main() {
	err := run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(args []string) error {
	fmt.Printf("hello world")
	return nil
}
