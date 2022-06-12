package main

import (
	"fmt"

	"github.com/AldieNightStar/flower"
	"github.com/AldieNightStar/golisper"
)

func main() {
	scope, err := flower.LoadFromFile("file.lsp")
	if err != nil {
		fmt.Println(err)
		return
	}
	addApi(scope)
	_, err = scope.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func addApi(s *flower.Scope) {
	s.Api["print"] = func(s *flower.Scope, args []*golisper.Value) (any, error) {
		elems, err := s.EvalArrayValues(args)
		if err != nil {
			return nil, err
		}
		fmt.Println(elems...)
		return nil, nil
	}
}
