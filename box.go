package flower

type builtinBox struct {
	value any
}

func (b *builtinBox) GetValue(name string) any {
	if name == "value" {
		return b.value
	}
	return nil
}

func (b *builtinBox) SetValue(name string, val any) bool {
	if name == "value" {
		b.value = val
		return true
	}
	return false
}
