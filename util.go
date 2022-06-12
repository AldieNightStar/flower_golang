package flower

import "github.com/AldieNightStar/golisper"

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

func utilReadTagValues(val *golisper.Value, name string) []*golisper.Value {
	if val.Type != golisper.TYPE_TAG {
		return nil
	}
	if val.TagVal.Name != name {
		return nil
	}
	return val.TagVal.Values
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

func utilValuesToTags(vals []*golisper.Value) []*golisper.Tag {
	tags := make([]*golisper.Tag, 0, len(vals))
	for _, t := range vals {
		if t.Type == golisper.TYPE_TAG {
			tags = append(tags, t.TagVal)
		}
	}
	return tags
}

func utilReadTag(val *golisper.Value) *golisper.Tag {
	if val == nil || val.Type != golisper.TYPE_TAG {
		return nil
	}
	return val.TagVal
}

func utilReadTagWithName(name string, val *golisper.Value) *golisper.Tag {
	tag := utilReadTag(val)
	if tag == nil || tag.Name != name {
		return nil
	}
	return tag
}

func utilCodeFuncToSFunc(parent *Scope, f *codeFunction) SFunc {
	return func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < len(f.aliases) {
			return nil, errNotEnoughArgs(s.LastLine, "call function", len(f.aliases), len(args))
		}
		with := make(map[string]any)
		for id, alias := range f.aliases {
			evaledArg, err := s.Eval(args[id])
			if err != nil {
				return nil, err
			}
			with[alias] = evaledArg
		}
		scope := f.block.Load(parent, with)
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
