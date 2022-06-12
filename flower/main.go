package main

import (
	"fmt"

	"github.com/AldieNightStar/flower"
)

func main() {
	scope, err := flower.LoadFromFile("file.lsp")
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
