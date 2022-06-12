package flower

import "github.com/AldieNightStar/golisper"

func builtinBlocks(s *Scope) {
	s.Api["do"] = func(s *Scope, args []*golisper.Value) (any, error) {
		return newBlock(utilValuesToTags(args)), nil
	}
	s.Api["call"] = func(s *Scope, args []*golisper.Value) (any, error) {
		// (call block (with k val) (with k val))
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "call", 1, 0)
		}
		block, err := EvalCast[*codeBlock]("call", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		with, err := utilReadAllWithTags(s, args[1:])
		if err != nil {
			return nil, err
		}
		blockScope := block.Load(s, with)
		return blockScope.Run()
	}
}
