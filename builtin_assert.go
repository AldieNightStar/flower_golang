package flower

import "github.com/AldieNightStar/golisper"

func builtinAssert(s *Scope) {
	s.Memory["assert"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "assert", 1, 0)
		}
		b, err := EvalCast("assert", s, args[0], false)
		if err != nil {
			return nil, err
		}
		if !b {
			if len(args) > 1 {
				message, err := EvalCast("assert", s, args[1], "")
				if err != nil {
					return nil, err
				}
				return nil, newErrLineName(s.LastLine, "assert", "Assertion fail. Message: "+message)
			} else {
				return nil, newErrLineName(s.LastLine, "assert", "Assertion fail")
			}
		}
		return nil, nil
	})
}
