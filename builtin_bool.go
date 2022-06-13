package flower

import (
	"github.com/AldieNightStar/golisper"
)

func builtinBool(s *Scope) {
	s.Memory["true"] = true
	s.Memory["false"] = false
	s.Memory["nil"] = nil
	s.Memory["eq"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
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
	})
	s.Memory["less"] = builtinBoolNumberOp(func(a, b float64) bool { return a < b })
	s.Memory["greater"] = builtinBoolNumberOp(func(a, b float64) bool { return a > b })
	s.Memory["less-eq"] = builtinBoolNumberOp(func(a, b float64) bool { return a <= b })
	s.Memory["greater-eq"] = builtinBoolNumberOp(func(a, b float64) bool { return a >= b })
	s.Memory["not"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "not operation", 1, 0)
		}
		b, err := EvalCast("not operation", s, args[0], false)
		if err != nil {
			return nil, err
		}
		return !b, nil
	})
	s.Memory["and"] = builtinBoolOp(func(a, b bool) bool { return a && b })
	s.Memory["or"] = builtinBoolOp(func(a, b bool) bool { return a || b })
	s.Memory["isnull"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "not operation", 1, 0)
		}
		val, err := s.Eval(args[0])
		if err != nil {
			return nil, err
		}
		return val != nil, nil
	})
	s.Memory["maybe"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 3 {
			return nil, errNotEnoughArgs(s.LastLine, "not operation", 3, len(args))
		}
		b, err := EvalCast("maybe operator", s, args[0], false)
		if err != nil {
			return nil, err
		}
		val1, err := s.Eval(args[1])
		if err != nil {
			return nil, err
		}
		val2, err := s.Eval(args[2])
		if err != nil {
			return nil, err
		}
		if b {
			return val1, nil
		}
		return val2, nil
	})
	s.Memory["is-nil"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "isnull", 1, 0)
		}
		val, err := s.Eval(args[0])
		if err != nil {
			return nil, err
		}
		return val == nil, nil
	})
	s.Memory["nil-def"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "isnull", 2, len(args))
		}
		val1, err := s.Eval(args[0])
		if err != nil {
			return nil, err
		}
		val2, err := s.Eval(args[1])
		if err != nil {
			return nil, err
		}
		if val1 == nil {
			return val2, nil
		}
		return val1, nil
	})
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
