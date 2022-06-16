package flower

import "sync"

type builtinBox struct {
	value any
	mut   *sync.Mutex
}

func (b *builtinBox) GetValue(name string) any {
	b.mut.Lock()
	defer b.mut.Unlock()
	if name == "value" {
		return b.value
	}
	return nil
}

func (b *builtinBox) SetValue(name string, val any) bool {
	b.mut.Lock()
	defer b.mut.Unlock()
	if name == "value" {
		b.value = val
		return true
	}
	return false
}

func (b *builtinBox) Type() string {
	return "box"
}
