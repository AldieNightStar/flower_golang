package flower

type builtinList struct {
	list []any
}

func (l *builtinList) Add(elem any) {
	l.list = append(l.list, elem)
}

func (l *builtinList) Get(id int) any {
	if id < 0 || id >= len(l.list) {
		return nil
	}
	return l.list[id]
}

func (l *builtinList) Set(id int, val any) bool {
	if id < 0 || id >= len(l.list) {
		return false
	}
	l.list[id] = val
	return true
}

func (l *builtinList) Iteration() builtinIteration {
	return &builtinArrayIteration{l.list, 0}
}
