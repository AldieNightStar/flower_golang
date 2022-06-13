package flower

import (
	"github.com/AldieNightStar/golisper"
)

func builtinIter(s *Scope) {
	s.Memory["of"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		return &builtinIteratorOfArgs{s, args}, nil
	})
	s.Memory["range"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "range", 2, len(args))
		}
		start, err := EvalCast[float64]("iterate", s, args[0], 0)
		if err != nil {
			return nil, err
		}
		end, err := EvalCast[float64]("iterate", s, args[1], 0)
		if err != nil {
			return nil, err
		}
		return &builtinRangeIterator{int(start), int(end)}, nil
	})
	s.Memory["infinite"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		return builtinForeverIterator(0), nil
	})
	s.Memory["iterate"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		// (iterate iterator item block)
		if len(args) < 3 {
			return nil, errNotEnoughArgs(s.LastLine, "iterate", 3, len(args))
		}
		iter, err := EvalCast[builtinIterator]("iterate", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		alias := utilReadEtcString(args[1])
		block, err := EvalCast[*codeBlock]("iterate", s, args[2], nil)
		if err != nil {
			return nil, err
		}
		iteration := iter.Iteration()
		toBreak := false
		scope := block.Load(s, nil)
		scope.Memory["break"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
			toBreak = true
			scope.Pos = 0xFFFFFFFF
			s.Pos = 0xFFFFFFFF
			s.IsEnded = true
			return nil, nil
		})
		scope.Memory["continue"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
			scope.Pos = 0xFFFFFFFF
			s.Pos = 0xFFFFFFFF
			return nil, nil
		})
		for {
			item, err := iteration.Iterate()
			if err != nil {
				return nil, err
			}
			if item == nil {
				break
			}
			scope.Memory[alias] = item
			_, err = scope.Run()
			if err != nil {
				return nil, err
			}
			if toBreak {
				break
			}
		}
		return nil, nil
	})
	s.Memory["iteration"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "iteration", 1, 0)
		}
		iter, err := EvalCast[builtinIterator]("iteration", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		return iter.Iteration(), nil
	})
	s.Memory["next"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "next", 1, 0)
		}
		iter, err := EvalCast[builtinIteration]("next", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		return iter.Iterate()
	})
	s.Memory["next-all"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "next", 1, 0)
		}
		iter, err := EvalCast[builtinIteration]("next", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		arr := &builtinList{make([]any, 0, 32)}
		for {
			elem, err := iter.Iterate()
			if err != nil {
				return nil, err
			}
			if elem == nil {
				break
			}
			arr.list = append(arr.list, elem)
		}
		return arr, nil
	})
}
