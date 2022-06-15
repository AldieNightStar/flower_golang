package flower

import (
	"github.com/AldieNightStar/golisper"
)

func builtinErrors(s *Scope) {
	s.Memory["error"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "safe", 1, 0)
		}
		errString, err := EvalCast("safe", s, args[0], "")
		if err != nil {
			return nil, err
		}
		return nil, newErrLineName(s.LastLine, "error", errString)
	})
	s.Memory["safe"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "safe", 1, 0)
		}
		block, err := EvalCast[*codeBlock]("safe", s, args[0], nil)
		if err != nil {
			return nil, newErrLineName(s.LastLine, "safe", err.Error())
		}
		scope := block.Load(block.scope, nil)
		builtinAddReturn(scope)
		_, err = scope.Run()
		if err != nil {
			return err, nil
		}
		return nil, nil
	})
}
