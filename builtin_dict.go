package flower

import "github.com/AldieNightStar/golisper"

func builtinDict(s *Scope) {
	s.Api["dict"] = func(s *Scope, args []*golisper.Value) (any, error) {
		evaled, err := s.EvalArrayValues(args)
		if err != nil {
			return nil, err
		}
		dict := utilCollectKeyValsToMap(evaled)
		return &builtinDictStruct{dict}, nil
	}
	s.Api["dict-get"] = func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "dict-get", 2, len(args))
		}
		dict, err := EvalCast[*builtinDictStruct]("dict-get", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		name, err := EvalCast("dict-get", s, args[1], "")
		if err != nil {
			return nil, err
		}
		item, ok := dict.m[name]
		if !ok {
			if len(args) > 2 {
				defVal, err := s.Eval(args[2])
				if err != nil {
					return nil, err
				}
				return defVal, nil
			}
			return nil, newErrLineName(s.LastLine, "dict-get", "Key '"+name+"' not found in dict")
		}
		return item, nil
	}
	s.Api["dict-set"] = func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 3 {
			return nil, errNotEnoughArgs(s.LastLine, "dict-set", 3, len(args))
		}
		dict, err := EvalCast[*builtinDictStruct]("dict-set", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		name, err := EvalCast("dict-set", s, args[1], "")
		if err != nil {
			return nil, err
		}
		val, err := s.Eval(args[2])
		if err != nil {
			return nil, err
		}
		dict.m[name] = val
		return val, nil
	}
	s.Api["dict-len"] = func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "dict-set", 1, 0)
		}
		dict, err := EvalCast[*builtinDictStruct]("dict-set", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		return float64(len(dict.m)), nil
	}
	s.Api["dict-keys"] = func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "dict-set", 1, 0)
		}
		dict, err := EvalCast[*builtinDictStruct]("dict-set", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		return &builtinDictKeysIterator{dict: dict}, nil
	}
}
