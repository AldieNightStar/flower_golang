package flower

import (
	"fmt"

	"github.com/AldieNightStar/golisper"
)

var builtins = (func() *Scope {
	scope := NewScope(nil, 0, nil)
	// Register code
	builtinMath(scope)
	builtinIter(scope)
	builtinBlocks(scope)
	builtinMem(scope)
	builtinBool(scope)
	builtinLoopIf(scope)
	builtinTime(scope)

	// Return command
	scope.Api["return"] = func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "return", 1, 0)
		}
		var err error
		s.ReturnVal, err = s.Eval(args[0])
		if err != nil {
			return nil, err
		}
		s.WillReturn = true
		return nil, nil
	}
	// Print command
	scope.Api["print"] = func(s *Scope, args []*golisper.Value) (any, error) {
		elems, err := s.EvalArrayValues(args)
		if err != nil {
			return nil, err
		}
		fmt.Println(elems...)
		return nil, nil
	}
	// End
	return scope
})()
