package flower

import "github.com/AldieNightStar/golisper"

type keyval struct {
	key string
	val any
}

type builtinExtends struct {
	dict *builtinDictStruct
}

func builtinKeyVal(s *Scope) {
	s.Memory["with"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "with", 2, len(args))
		}
		key, err := EvalCast("with", s, args[0], "")
		if err != nil {
			return nil, err
		}
		val, err := s.Eval(args[1])
		if err != nil {
			return nil, err
		}
		return &keyval{key, val}, nil
	})
	s.Memory["extends"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "extends", 1, 0)
		}
		dict, err := EvalCast[*builtinDictStruct]("extends", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		return &builtinExtends{dict}, nil
	})
}
