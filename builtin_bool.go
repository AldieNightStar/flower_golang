package flower

import "github.com/AldieNightStar/golisper"

func builtinBool(s *Scope) {
	s.Api["eq"] = func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "eq", 2, len(args))
		}
		o1, err := s.Eval(args[0])
		if err != nil {
			return nil, err
		}
		o2, err := s.Eval(args[1])
		if err != nil {
			return nil, err
		}
		return o1 == o2, nil
	}
	s.Api["less"] = builtinBoolNumberOp(func(a, b float64) bool { return a < b })
	s.Api["greater"] = builtinBoolNumberOp(func(a, b float64) bool { return a > b })
	s.Api["less-eq"] = builtinBoolNumberOp(func(a, b float64) bool { return a <= b })
	s.Api["greater-eq"] = builtinBoolNumberOp(func(a, b float64) bool { return a >= b })
	s.Api["not"] = func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "not operation", 1, 0)
		}
		b, err := EvalCast("not operation", s, args[0], false)
		if err != nil {
			return nil, err
		}
		return !b, nil
	}
	s.Api["and"] = builtinBoolOp(func(a, b bool) bool { return a && b })
	s.Api["or"] = builtinBoolOp(func(a, b bool) bool { return a || b })
}

func builtinBoolNumberOp(f func(a, b float64) bool) SFunc {
	return func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "comparing numbers", 2, len(args))
		}
		f1, err := EvalCast[float64]("comparing numbers", s, args[0], 0)
		if err != nil {
			return nil, err
		}
		f2, err := EvalCast[float64]("comparing numbers", s, args[1], 0)
		if err != nil {
			return nil, err
		}
		return f(f1, f2), nil
	}
}

func builtinBoolOp(f func(a, b bool) bool) SFunc {
	return func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "comparing numbers", 2, len(args))
		}
		b1, err := EvalCast[bool]("comparing numbers", s, args[0], false)
		if err != nil {
			return nil, err
		}
		b2, err := EvalCast[bool]("comparing numbers", s, args[1], false)
		if err != nil {
			return nil, err
		}
		return f(b1, b2), nil
	}
}
