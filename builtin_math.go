package flower

import (
	"github.com/AldieNightStar/golisper"
)

func builtinMath(s *Scope) {
	s.Api["add"] = builtinMathOp(func(a, b float64) float64 { return a + b })
	s.Api["sub"] = builtinMathOp(func(a, b float64) float64 { return a - b })
	s.Api["mul"] = builtinMathOp(func(a, b float64) float64 { return a * b })
	s.Api["div"] = builtinMathOp(func(a, b float64) float64 { return a / b })
	s.Api["mod"] = builtinMathOp(func(a, b float64) float64 { return float64(int(a) % int(b)) })
}

func builtinMathOp(f func(a, b float64) float64) SFunc {
	return func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return 0, nil
		}
		e1, err := s.Eval(args[0])
		if err != nil {
			return nil, err
		}
		e2, err := s.Eval(args[1])
		if err != nil {
			return nil, err
		}
		f1, ok := e1.(float64)
		if !ok {
			return 0, nil
		}
		f2, ok := e2.(float64)
		if !ok {
			return 0, nil
		}
		return f(f1, f2), nil
	}
}
