package flower

import (
	"strings"

	"github.com/AldieNightStar/golisper"
)

func builtinMem(s *Scope) {
	s.Memory["set"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "set", 2, len(args))
		}
		name := utilReadEtcString(args[0])
		newVal, err := s.Eval(args[1])
		if err != nil {
			return nil, err
		}
		if path := utilReadPathVariableName(name); path != nil {
			// If name has dots
			// Get the dictionary and write value to it
			// [!] last name value is key for the dictionary
			key := path[len(path)-1]
			path = path[0 : len(path)-1]
			val, err := utilEvalPathVariable(s, path)
			if err != nil {
				return nil, err
			}
			if dict, ok := val.(builtinValuesSetter); ok {
				dict.SetValue(key, newVal)
				return nil, nil
			}
			return nil, newErrLineName(s.LastLine, "set", "Can't set value inside non-dict")
		} else {
			s.Memory[name] = newVal
		}
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
	s.Memory["as"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "with", 2, len(args))
		}
		val, err := s.Eval(args[0])
		if err != nil {
			return nil, err
		}
		block, err := EvalCast[*codeBlock]("with", s, args[1], nil)
		if err != nil {
			return nil, err
		}
		scope := block.Load(block.scope, map[string]any{"it": val})
		builtinAddReturn(scope)
		return scope.Run()
	})
}
