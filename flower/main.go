package main

import (
	"fmt"
	"os"

	"github.com/AldieNightStar/flower"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Enter file to run:\n\tflower file.lsp")
		return
	}
	scope, err := flower.LoadFromFile(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	scope.AllowFileAccess()
	scope.AllowHTTP()
	_, err = scope.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
