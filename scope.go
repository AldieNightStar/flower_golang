package flower

import (
	"fmt"

	"github.com/AldieNightStar/golisper"
)

type SFunc func(s *Scope, args []*golisper.Value) (any, error)

type Scope struct {
	Code       []*golisper.Tag
	Pos        int
	Parent     *Scope
	Api        map[string]SFunc
	Memory     map[string]any
	ReturnVal  any
	WillReturn bool
	IsEnded    bool
	LastLine   int
}

func NewScopeWithBuiltIns(code []*golisper.Tag, Pos int) *Scope {
	return NewScope(code, Pos, builtins)
}

func NewScope(Code []*golisper.Tag, Pos int, Parent *Scope) *Scope {
	return &Scope{
		Code:       Code,
		Pos:        Pos,
		Api:        make(map[string]SFunc, 64),
		Parent:     Parent,
		Memory:     make(map[string]any, 64),
		ReturnVal:  nil,
		WillReturn: false,
		IsEnded:    false,
	}
}

func (s *Scope) Next() any {
	if s.Pos >= len(s.Code) {
		return nil
	}
	res := s.Code[s.Pos]
	s.LastLine = res.Line
	s.Pos += 1
	return res
}

func (s *Scope) GetApiFunc(name string) SFunc {
	f, ok := s.Api[name]
	if !ok {
		if s.Parent == nil {
			return nil
		}
		return s.Parent.GetApiFunc(name)
	}
	return f
}

func (s *Scope) GetVariableValue(name string) any {
	v, ok := s.Memory[name]
	if !ok {
		if s.Parent == nil {
			return nil
		}
		return s.Parent.GetVariableValue(name)
	}
	return v
}

func (s *Scope) Eval(tok any) (any, error) {
	if tag, tagOk := tok.(*golisper.Tag); tagOk {
		tagName := tag.Name
		f := s.GetApiFunc(tagName)
		if f == nil {
			return nil, fmt.Errorf("api function '%s' is not exist. Line: %d", tagName, tag.Line)
		}
		return f(s, tag.Values)
	} else if val, valOk := tok.(*golisper.Value); valOk {
		if val.Type == golisper.TYPE_ETC_STRING {
			return s.GetVariableValue(val.StringVal), nil
		} else if val.Type == golisper.TYPE_STRING {
			return val.StringVal, nil
		} else if val.Type == golisper.TYPE_NUMBER {
			return val.NumberVal, nil
		} else if val.Type == golisper.TYPE_TAG {
			return s.Eval(val.TagVal)
		} else {
			// TODO: May be better to make error
			return nil, nil
		}
	}
	return tok, nil
}

func (s *Scope) EvalArrayValues(arr []*golisper.Value) ([]any, error) {
	res := make([]any, 0, len(arr))
	for _, elem := range arr {
		val, err := s.Eval(elem)
		if err != nil {
			return nil, err
		}
		res = append(res, val)
	}
	return res, nil
}

func (s *Scope) Step() error {
	tok := s.Next()
	if tok == nil {
		s.IsEnded = true
		return nil
	}
	_, err := s.Eval(tok)
	if err != nil {
		return err
	}
	return nil
}

func (s *Scope) Run() (any, error) {
	for {
		err := s.Step()
		if err != nil {
			return nil, err
		}
		if s.WillReturn {
			res := s.ReturnVal
			s.WillReturn = false
			s.ReturnVal = nil
			return res, nil
		}
		if s.IsEnded {
			return nil, nil
		}
	}
}

func (s *Scope) LocalScope(vals []*golisper.Value) *Scope {
	return NewScope(utilValuesToTags(vals), 0, s)
}
