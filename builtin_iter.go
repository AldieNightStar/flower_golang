package flower

import (
	"github.com/AldieNightStar/golisper"
)

type builtinIterConstructor struct {
	values []*golisper.Value
}

func newBuiltinIterConstructor(values []*golisper.Value) *builtinIterConstructor {
	return &builtinIterConstructor{values: values}
}

func (i *builtinIterConstructor) Iterate() *builtinIterIteration {
	return &builtinIterIteration{
		values: i.values,
		ptr:    0,
	}
}

type builtinIterIteration struct {
	values []*golisper.Value
	ptr    int
}

func (i *builtinIterIteration) Iterate() *golisper.Value {
	if i.ptr < len(i.values) {
		res := i.values[i.ptr]
		i.ptr += 1
		return res
	}
	return nil
}

func builtinIter(s *Scope) {
	s.Api["of"] = func(s *Scope, args []*golisper.Value) (any, error) {
		return newBuiltinIterConstructor(args), nil
	}
	// s.Api["iterate"] = func(s *Scope, args []*golisper.Value) (any, error) {
	// 	// (iterate iterator item block)
	// 	if len(args) < 3 {
	// 		return nil, errNotEnoughArgs(s.LastLine, "iterate", 3, len(args))
	// 	}
	// 	iter, err := EvalCast[*builtinIterConstructor]("iterate", s, args[0], nil)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	alias := utilReadEtcString(args[1])
	// 	block := utilReadBlock(args[2], "do")
	// 	if block == nil {

	// 	}
	// }
}
