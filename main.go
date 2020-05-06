package main

import (
	"fmt"
	"os"
)

var (
	version string = `¯\_(ツ)_/¯`
	commit  string = `¯\_(ツ)_/¯`
	date    string = `¯\_(ツ)_/¯`
	builtBy string = `¯\_(ツ)_/¯`
)

func main() {
	err := run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(args []string) error {
	fmt.Println("Hey! You!")
	fmt.Printf("I am running go-release-example!\n\nVersion:\t%s\nCommit:\t\t%s\nDate:\t\t%s\nBuilt By:\t%s\n", version, commit, date, builtBy)
	return nil
}
