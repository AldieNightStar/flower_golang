package flower

import "github.com/AldieNightStar/golisper"

func builtinType(s *Scope) {
	s.Memory["type"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "type", 1, 0)
		}
		v, err := s.Eval(args[0])
		if err != nil {
			return nil, err
		}
		return getTypeOf(v), nil
	})
}
