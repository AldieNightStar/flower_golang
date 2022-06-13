package flower

import (
	"strings"

	"github.com/AldieNightStar/golisper"
)

type builtinCell struct {
	val any
}

func builtinMem(s *Scope) {
	s.Memory["set"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "set", 2, len(args))
		}
		name := utilReadEtcString(args[0])
		if strings.Contains(name, ".") {
			return nil, newErrLineName(s.LastLine, "set", "Path variable not allowed here")
		}
		val, err := s.Eval(args[1])
		if err != nil {
			return nil, err
		}
		s.Memory[name] = val
		return nil, nil
	})
	s.Memory["cell"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		var val any
		var err error
		if len(args) > 0 {
			val, err = s.Eval(args[0])
			if err != nil {
				return nil, err
			}
		}
		return &builtinCell{val}, nil
	})
	s.Memory["cell-get"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "set", 1, 0)
		}
		cell, err := EvalCast[*builtinCell]("cell-set", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		return cell.val, nil
	})
	s.Memory["cell-set"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "set", 2, len(args))
		}
		cell, err := EvalCast[*builtinCell]("cell-set", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		val, err := s.Eval(args[1])
		if err != nil {
			return nil, err
		}
		cell.val = val
		return nil, nil
	})
	s.Memory["require"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "set", 1, 0)
		}
		absent := make([]string, 0, 8)
		for _, arg := range args {
			etc := utilReadEtcString(arg)
			if etc == "" {
				return nil, newErrLineName(s.LastLine, "require", "Should be only constant names")
			}
			_, ok := s.Memory[etc]
			if !ok {
				absent = append(absent, etc)
			}
		}
		if len(absent) > 0 {
			return nil, newErrLineName(
				s.LastLine, "require", "Variables are absent: "+strings.Join(absent, ", "),
			)
		}
		return nil, nil
	})
}
