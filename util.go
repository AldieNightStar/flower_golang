package flower

import "github.com/AldieNightStar/golisper"

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

// returns key, value, nil if success.
// Can also return "", nil, nil if tag is not "with"
func utilEvalTagWith(scope *Scope, v *golisper.Value) (key string, val any, err error) {
	vals := utilReadTagValues(v, "with")
	if vals == nil {
		return "", nil, nil
	}
	if len(vals) < 2 {
		return "", nil, errNotEnoughArgs(scope.LastLine, "with", 2, len(vals))
	}
	key = utilReadEtcString(vals[0])
	if key == "" {
		return "", nil, newErrLineName(scope.LastLine, "with", "Key value is empty or not set!")
	}
	val, err = scope.Eval(vals[1])
	if err != nil {
		return "", nil, err
	}
	return key, val, nil
}

func utilReadAllWithTags(scope *Scope, vals []*golisper.Value) (map[string]any, error) {
	with := make(map[string]any)
	for _, a := range vals {
		k, val, err := utilEvalTagWith(scope, a)
		if k == "" && val == nil {
			continue
		}
		if err != nil {
			return nil, err
		}
		with[k] = val
	}
	return with, nil
}
