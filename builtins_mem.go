package flower

import "github.com/AldieNightStar/golisper"

func builtinMem(s *Scope) {
	s.Api["set"] = func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "set", 2, len(args))
		}
		name := utilReadEtcString(args[0])
		val, err := s.Eval(args[1])
		if err != nil {
			return nil, err
		}
		s.Memory[name] = val
		return nil, nil
	}
}
