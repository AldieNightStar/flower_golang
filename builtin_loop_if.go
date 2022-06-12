package flower

import "github.com/AldieNightStar/golisper"

func builtinLoopIf(s *Scope) {
	s.Api["if"] = func(s *Scope, args []*golisper.Value) (any, error) {
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
			return block.Load(s, nil).Run()
		} else if elseBlock != nil {
			return elseBlock.Load(s, nil).Run()
		}
		// TODO: May be some error here?
		return nil, nil
	}
	s.Api["repeat"] = func(s *Scope, args []*golisper.Value) (any, error) {
		// (repeat 10 i (do ...))
		if len(args) < 3 {
			return nil, errNotEnoughArgs(s.LastLine, "repeat", 3, len(args))
		}
		numF, err := EvalCast[float64]("repeat", s, args[0], 0)
		if err != nil {
			return nil, err
		}
		num := int(numF)
		alias := utilReadEtcString(args[1])
		block, err := EvalCast[*codeBlock]("repeat", s, args[2], nil)
		if err != nil {
			return nil, err
		}
		if num < 0 {
			return nil, newErrLineName(s.LastLine, "repeat", "Iteration count is less than 0")
		}
		scope := block.Load(s, map[string]any{alias: float64(0)})
		toBreak := false
		scope.Api["break"] = func(s *Scope, args []*golisper.Value) (any, error) {
			s.Pos = 0xFFFFFFFF
			toBreak = true
			return nil, nil
		}
		for {
			f := scope.Memory[alias].(float64)
			if int(f) >= num {
				break
			}
			scope.Pos = 0
			_, err := scope.Run()
			if err != nil {
				return nil, err
			}
			if toBreak {
				break
			}
			scope.Memory[alias] = f + 1
		}
		return nil, nil
	}
}
