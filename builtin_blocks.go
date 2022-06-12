package flower

import "github.com/AldieNightStar/golisper"

type codeFunction struct {
	aliases []string
	block   *codeBlock
}

func builtinBlocks(s *Scope) {
	s.Memory["do"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		return newBlock(utilValuesToTags(args)), nil
	})
	s.Memory["call"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		// (call block (with k val) (with k val))
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "call", 1, 0)
		}
		block, err := EvalCast[*codeBlock]("call", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		evaledArgs, err := s.EvalArrayValues(args[1:])
		if err != nil {
			return nil, err
		}
		blockScope := block.Load(s, utilCollectKeyValsToMap(evaledArgs))
		return blockScope.Run()
	})
	s.Memory["def"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
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
		codeFunc := &codeFunction{
			aliases: aliases,
			block:   block,
		}
		return utilCodeFuncToSFunc(codeFunc), nil
	})
}
