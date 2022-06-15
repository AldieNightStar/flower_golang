package flower

import "github.com/AldieNightStar/golisper"

type codeBlock struct {
	code  []*golisper.Value
	scope *Scope
}

func (b *codeBlock) Type() string {
	return "block"
}

func newBlock(scope *Scope, code []*golisper.Value) *codeBlock {
	return &codeBlock{code, scope}
}

func (b *codeBlock) Load(parent *Scope, with map[string]any) *Scope {
	scope := NewScope(b.code, 0, parent)
	for k, v := range with {
		scope.Memory[k] = v
	}
	return scope
}
