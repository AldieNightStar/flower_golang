package flower

import "github.com/AldieNightStar/golisper"

func builtinsList(s *Scope) {
	s.Api["list"] = func(s *Scope, args []*golisper.Value) (any, error) {
		arr := make([]any, 0, 32)
		if len(args) > 0 {
			iter, err := EvalCast[builtinIterator]("list", s, args[0], nil)
			if err != nil {
				return nil, err
			}
			iteration := iter.Iteration()
			for {
				elem, err := iteration.Iterate()
				if err != nil {
					return nil, err
				}
				if elem == nil {
					break
				}
				arr = append(arr, elem)
			}
		}
		return &builtinList{arr}, nil
	}
	s.Api["list-get"] = func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "list-get", 2, len(args))
		}
		list, err := EvalCast[*builtinList]("list-get", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		idF, err := EvalCast[float64]("list-get", s, args[1], 0)
		if err != nil {
			return nil, err
		}
		id := int(idF)
		return list.Get(id), nil
	}
	s.Api["list-set"] = func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 3 {
			return nil, errNotEnoughArgs(s.LastLine, "list-set", 3, len(args))
		}
		list, err := EvalCast[*builtinList]("list-set", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		idF, err := EvalCast[float64]("list-set", s, args[1], 0)
		if err != nil {
			return nil, err
		}
		id := int(idF)
		newVal, err := EvalCast[any]("list-set", s, args[2], nil)
		if err != nil {
			return nil, err
		}
		return list.Set(id, newVal), nil
	}
	s.Api["list-add"] = func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "list-add", 2, len(args))
		}
		list, err := EvalCast[*builtinList]("list-add", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		newVal, err := EvalCast[any]("list-add", s, args[1], nil)
		if err != nil {
			return nil, err
		}
		list.Add(newVal)
		return newVal, nil
	}
	s.Api["list-len"] = func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "list-len", 1, 0)
		}
		list, err := EvalCast[*builtinList]("list-len", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		return float64(len(list.list)), nil
	}
}
