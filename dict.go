package flower

type builtinDictStruct struct {
	m map[string]any
}

func newBuiltinDict() *builtinDictStruct {
	return &builtinDictStruct{make(map[string]any)}
}
