package flower

import (
	"sync"

	"github.com/AldieNightStar/golisper"
)

func builtinThreading(s *Scope) {
	s.Memory["thread"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "thread", 1, 0)
		}
		block, err := EvalCast[*codeBlock]("thread", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		scope := block.Load(block.scope, nil)
		builtinAddReturn(scope)
		go scope.Run()
		return nil, nil
	})
	s.Memory["mutex"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		return &sync.Mutex{}, nil
	})
	s.Memory["lock"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "lock", 2, len(args))
		}
		mut, err := EvalCast[*sync.Mutex]("lock", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		block, err := EvalCast[*codeBlock]("lock", s, args[1], nil)
		if err != nil {
			return nil, err
		}
		scope := block.Load(block.scope, nil)
		builtinAddReturn(scope)
		mut.Lock()
		defer mut.Unlock()
		return scope.Run()
	})
}
