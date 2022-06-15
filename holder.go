package flower

type builtinValuesGetter interface {
	GetValue(name string) any
}

type builtinValuesSetter interface {
	SetValue(name string, val any) bool
}

type builtinValueLen interface {
	Len() int
}

type builtinKeysHolder interface {
	Keys() []string
}
