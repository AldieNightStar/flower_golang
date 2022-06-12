package flower

import "github.com/AldieNightStar/golisper"

type codeFunction struct {
	aliases []string
	block   *codeBlock
}

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
	s.Api["def"] = func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "def", 1, 0)
		}
		aliases := make([]string, 0, 8)
		var block *codeBlock
		var err error
		for _, arg := range args {
			alias := utilReadEtcString(arg)
			if alias != "" {
				aliases = append(aliases, alias)
				continue
			}
			block, err = EvalCast[*codeBlock]("def", s, arg, nil)
			if err != nil {
				return nil, err
			}
		}
		return &codeFunction{
			aliases: aliases,
			block:   block,
		}, nil
	}
}
