package flower

import "github.com/AldieNightStar/golisper"

func builtinLoopIf(s *Scope) {
	s.Memory["if"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		// (if (eq 2 2) (do ...))
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "if", 2, len(args))
		}
		b, err := EvalCast("if", s, args[0], false)
		if err != nil {
			return nil, err
		}
		block, err := EvalCast[*codeBlock]("if", s, args[1], nil)
		if err != nil {
			return nil, err
		}
		var elseBlock *codeBlock = nil
		if len(args) > 2 {
			elseBlock, err = EvalCast[*codeBlock]("if", s, args[2], nil)
			if err != nil {
				return nil, err
			}
		}
		if b {
			return block.Load(block.scope, nil).Run()
		} else if elseBlock != nil {
			return elseBlock.Load(elseBlock.scope, nil).Run()
		}
		// TODO: May be some error here?
		return nil, nil
	})
}
