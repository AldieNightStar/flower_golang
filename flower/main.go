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
	s.Api["mul"] = func(s *flower.Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return 0, nil
		}
		evl, err := s.EvalArrayValues(args)
		if err != nil {
			return nil, err
		}
		a := evl[0].(float64)
		b := evl[1].(float64)
		return a * b, nil
	}
}
