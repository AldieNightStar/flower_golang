package flower

import (
	"sync"

	"github.com/AldieNightStar/golisper"
)

func builtinsList(s *Scope) {
	listDict := newBuitinDict()
	listDict.m["new"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		arr := make([]any, 0, 32)
		if len(args) > 0 {
			iter, err := EvalCast[builtinIterator]("list new", s, args[0], nil)
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
		return &builtinList{arr, &sync.Mutex{}}, nil
	})
	listDict.m["get"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "list get", 2, len(args))
		}
		list, err := EvalCast[*builtinList]("list get", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		idF, err := EvalCast[float64]("list get", s, args[1], 0)
		if err != nil {
			return nil, err
		}
		id := int(idF)
		return list.Get(id), nil
	})
	listDict.m["set"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 3 {
			return nil, errNotEnoughArgs(s.LastLine, "list set", 3, len(args))
		}
		list, err := EvalCast[*builtinList]("list set", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		idF, err := EvalCast[float64]("list set", s, args[1], 0)
		if err != nil {
			return nil, err
		}
		id := int(idF)
		newVal, err := EvalCast[any]("list set", s, args[2], nil)
		if err != nil {
			return nil, err
		}
		return list.Set(id, newVal), nil
	})
	listDict.m["add"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "list add", 2, len(args))
		}
		list, err := EvalCast[*builtinList]("list add", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		newVal, err := EvalCast[any]("list add", s, args[1], nil)
		if err != nil {
			return nil, err
		}
		list.Add(newVal)
		return newVal, nil
	})
	s.Memory["list"] = listDict

	stackDict := newBuitinDict()
	stackDict.m["new"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		stack := newStack[any](1024)
		if len(args) > 0 {
			iter, err := EvalCast[builtinIterator]("stack new", s, args[0], nil)
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
				stack.Push(elem)
			}
		}
		return stack, nil
	})
	stackDict.m["pop"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "stack pop", 1, 0)
		}
		stack, err := EvalCast[*stack[any]]("stack pop", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		val, _ := stack.Pop(nil)
		return val, nil
	})
	stackDict.m["push"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "stack push", 2, len(args))
		}
		stack, err := EvalCast[*stack[any]]("stack push", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		val, err := s.Eval(args[1])
		if err != nil {
			return nil, err
		}
		return stack.Push(val), nil
	})
	s.Memory["stack"] = stackDict

	s.Memory["len"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "len", 1, 0)
		}
		valLen, err := EvalCast[builtinValueLen]("len", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		return float64(valLen.Len()), nil
	})
}
