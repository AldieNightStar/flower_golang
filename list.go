package flower

import (
	"fmt"
	"strings"
	"sync"
)

type builtinList struct {
	list []any
	mut  *sync.Mutex
}

func (l *builtinList) Type() string {
	return "list"
}

func (l *builtinList) Add(elem any) {
	l.mut.Lock()
	defer l.mut.Unlock()
	l.list = append(l.list, elem)
}

func (l *builtinList) Get(id int) any {
	l.mut.Lock()
	defer l.mut.Unlock()
	if id < 0 || id >= len(l.list) {
		return nil
	}
	return l.list[id]
}

func (l *builtinList) Set(id int, val any) bool {
	l.mut.Lock()
	defer l.mut.Unlock()
	if id < 0 || id >= len(l.list) {
		return false
	}
	l.list[id] = val
	return true
}

func (l *builtinList) Iteration() builtinIteration {
	l.mut.Lock()
	defer l.mut.Unlock()
	return &builtinArrayIteration{l.list, 0}
}

func (l builtinList) String() string {
	arr := make([]string, 0, 8)
	for k, v := range l.list {
		arr = append(arr, fmt.Sprintf("[%d] = ", k)+fmt.Sprint(v))
	}
	return "LIST [" + strings.Join(arr, ", ") + "]"
}

func (l *builtinList) Len() int {
	l.mut.Lock()
	defer l.mut.Unlock()
	return len(l.list)
}
