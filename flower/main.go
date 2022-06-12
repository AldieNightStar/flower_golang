package main

import (
	"fmt"
	"os"

	"github.com/AldieNightStar/flower"
)

func main() {
	fileName := LoadNameFromArgs(true) // debug
	if fileName == "" {
		return
	}
	scope, err := flower.LoadFromFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = scope.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func LoadNameFromArgs(debug bool) string {
	if debug {
		return "file.lsp"
	}
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Enter file to run:\n\tflower file.lsp")
		return ""
	}
	return args[0]
}
