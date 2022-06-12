package flower

import "github.com/AldieNightStar/golisper"

var builtins = (func() *Scope {
	scope := NewScope(nil, 0, nil)
	// Register code
	builtinMath(scope)
	builtinIter(scope)
	builtinBlocks(scope)
	builtinMem(scope)

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
	// End
	return scope
})()
