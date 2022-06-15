package flower

import "fmt"

type builtinTypeGetter interface {
	Type() string
}

func getTypeOf(a any) string {
	if a == nil {
		return "nil"
	}
	if t, ok := a.(builtinTypeGetter); ok {
		return t.Type()
	}
	if _, ok := a.(string); ok {
		return "string"
	}
	if _, ok := a.(float64); ok {
		return "number"
	}
	if _, ok := a.(bool); ok {
		return "bool"
	}
	return fmt.Sprintf("%T", a)
}
