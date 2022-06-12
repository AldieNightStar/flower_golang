package flower

import (
	"github.com/AldieNightStar/golisper"
)

func builtinIter(s *Scope) {
	s.Api["of"] = func(s *Scope, args []*golisper.Value) (any, error) {
		return &builtinIteratorOfArgs{s, args}, nil
	}
	s.Api["range"] = func(s *Scope, args []*golisper.Value) (any, error) {
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
	}
	s.Api["iterate"] = func(s *Scope, args []*golisper.Value) (any, error) {
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
		scope.Api["break"] = func(s *Scope, args []*golisper.Value) (any, error) {
			toBreak = true
			s.Pos = 0xFFFFFFFF
			return nil, nil
		}
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
	}
}
