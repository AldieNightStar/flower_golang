package flower

import (
	"fmt"
	"strings"

	"github.com/AldieNightStar/golisper"
)

type builtinDictStruct struct {
	m     map[string]any
	super *builtinDictStruct
}

func (d builtinDictStruct) String() string {
	arr := make([]string, 0, 8)
	for k, v := range d.m {
		arr = append(arr, "["+k+"] = "+fmt.Sprint(v))
	}
	return "DICT [" + strings.Join(arr, ", ") + "]"
}

func newBuitinDict() *builtinDictStruct {
	return &builtinDictStruct{
		m:     make(map[string]any),
		super: nil,
	}
}

func (d *builtinDictStruct) GetValue(name string) any {
	val, ok := d.m[name]
	if !ok {
		// May be try to find in super dict
		if d.super != nil {
			return d.super.GetValue(name)
		}
		return nil
	}
	return val
}

func builtinDict(s *Scope) {
	d := newBuitinDict()
	d.m["new"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		evaled, err := s.EvalArrayValues(args)
		if err != nil {
			return nil, err
		}
		ext := utilFindExtendsTag(evaled)
		dict := utilCollectKeyValsToMap(evaled)
		dictSturct := &builtinDictStruct{dict, nil}
		if ext != nil {
			dictSturct.super = ext.dict
		}
		return dictSturct, nil
	})
	d.m["get"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "dict get", 2, len(args))
		}
		dict, err := EvalCast[*builtinDictStruct]("dict get", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		name, err := EvalCast("dict-get", s, args[1], "")
		if err != nil {
			return nil, err
		}
		item := dict.GetValue(name)
		if item == nil {
			if len(args) > 2 {
				defVal, err := s.Eval(args[2])
				if err != nil {
					return nil, err
				}
				return defVal, nil
			}
			return nil, newErrLineName(s.LastLine, "dict get", "Key '"+name+"' not found in dict")
		}
		return item, nil
	})
	d.m["set"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 3 {
			return nil, errNotEnoughArgs(s.LastLine, "dict set", 3, len(args))
		}
		dict, err := EvalCast[*builtinDictStruct]("dict set", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		name, err := EvalCast("dict set", s, args[1], "")
		if err != nil {
			return nil, err
		}
		val, err := s.Eval(args[2])
		if err != nil {
			return nil, err
		}
		dict.m[name] = val
		return val, nil
	})
	d.m["len"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "dict len", 1, 0)
		}
		dict, err := EvalCast[*builtinDictStruct]("dict len", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		return float64(len(dict.m)), nil
	})
	d.m["keys"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "dict keys", 1, 0)
		}
		dict, err := EvalCast[*builtinDictStruct]("dict keys", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		return &builtinDictKeysIterator{dict: dict}, nil
	})
	s.Memory["dict"] = d
}
