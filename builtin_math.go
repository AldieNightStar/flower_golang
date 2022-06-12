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
		f1, err := EvalCast[float64]("math operation", s, args[0], 0)
		if err != nil {
			return nil, err
		}
		f2, err := EvalCast[float64]("math operation", s, args[1], 0)
		if err != nil {
			return nil, err
		}
		return f(f1, f2), nil
	}
}
