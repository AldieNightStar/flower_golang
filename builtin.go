package flower

import (
	"github.com/AldieNightStar/golisper"
)

var builtinScope = (func() *Scope {
	scope := NewScope(nil, 0, nil)
	// Register code
	builtinMath(scope)
	builtinIter(scope)
	builtinBlocks(scope)
	builtinMem(scope)
	builtinBool(scope)
	builtinLoopIf(scope)
	builtinTime(scope)
	builtinDict(scope)
	builtinKeyVal(scope)
	builtinsList(scope)
	builtinString(scope)
	builtinAssert(scope)

	// Return command
	scope.Memory["return"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
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
	})
	return scope
})()
