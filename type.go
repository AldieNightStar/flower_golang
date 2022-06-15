package flower

import (
	"fmt"

	"github.com/AldieNightStar/golisper"
)

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
	if lispVal, ok := a.(*golisper.Value); ok {
		if lispVal.Type == golisper.TYPE_ETC_STRING {
			return "token:etc_string"
		} else if lispVal.Type == golisper.TYPE_NUMBER {
			return "token:number"
		} else if lispVal.Type == golisper.TYPE_STRING {
			return "token:string"
		} else if lispVal.Type == golisper.TYPE_TAG {
			return "token:tag"
		}
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
