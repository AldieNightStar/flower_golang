package flower

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/AldieNightStar/golisper"
)

func utilIgnoreHashBangAtStart(src string) string {
	if src[0:2] != "#!" {
		return src
	}
	for id, c := range src {
		if c == '\n' {
			return src[id+1:]
		}
	}
	return src
}

func utilReadEtcString(val *golisper.Value) string {
	if val == nil || val.Type != golisper.TYPE_ETC_STRING {
		return ""
	}
	return val.StringVal
}

func EvalCast[T any](commandName string, s *Scope, val any, def T) (T, error) {
	e, err := s.Eval(val)
	if err != nil {
		return def, err
	}
	if o, ok := e.(T); ok {
		return o, nil
	}
	return def, errWrongType(s.LastLine, commandName, val, def)
}

func EvalCastMust[T any](commandName string, s *Scope, val any, def T) T {
	e, err := s.Eval(val)
	if err != nil {
		panic(err)
	}
	if o, ok := e.(T); ok {
		return o
	}
	panic(errWrongType(s.LastLine, commandName, val, def))
}

func utilCodeFuncToSFunc(f *codeFunction) SFunc {
	return func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < len(f.aliases) {
			return nil, errNotEnoughArgs(s.LastLine, "call function", len(f.aliases), len(args))
		}
		with := make(map[string]any)
		for id, alias := range f.aliases {
			evaledArg, err := f.scope.Eval(args[id])
			if err != nil {
				return nil, err
			}
			with[alias] = evaledArg
		}
		scope := f.block.Load(f.scope, with)
		builtinAddReturn(scope)
		return scope.Run()
	}
}

func utilCollectKeyValsToMap(vals []any) map[string]any {
	mp := make(map[string]any)
	for _, val := range vals {
		if kv, kvOk := val.(*keyval); kvOk {
			mp[kv.key] = kv.val
		}
	}
	return mp
}

func utilFindExtendsTag(vals []any) *builtinExtends {
	for _, val := range vals {
		if e, eok := val.(*builtinExtends); eok {
			return e
		}
	}
	return nil
}

func utilReadPathVariableName(str string) []string {
	if !strings.Contains(str, ".") {
		return nil
	}
	return strings.Split(str, ".")
}

func utilEvalPathVariable(s *Scope, path []string) (any, error) {
	var holder builtinValuesGetter
	var val any
	for id, name := range path {
		if id == 0 {
			val = s.GetVariableValue(name)
		} else {
			val = holder.GetValue(name)
		}
		if id == len(path)-1 { // If last
			return val, nil
		}
		if val == nil {
			return nil, newErrLineName(s.LastLine, "variable path", "Variable path leads to nil in the middle: "+name)
		}
		if valHolder, valHolderOk := val.(builtinValuesGetter); valHolderOk {
			holder = valHolder
			continue
		}
		return nil, newErrLineName(s.LastLine, "variable path", "Variable path leads to a non dict/holder value in the middle: "+name)
	}
	return nil, nil
}

func utilReadFileAsString(name string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		return "", err
	}
	dat, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}
