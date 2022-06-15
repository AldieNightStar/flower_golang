package flower

import "github.com/AldieNightStar/golisper"

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
}
