package flower

type stack[T any] struct {
	Data []T
	Ptr  int
}

func newStack[T any](size int) *stack[T] {
	return &stack[T]{Data: make([]T, size), Ptr: 0}
}

func (s *stack[T]) Push(dat T) bool {
	if s.Ptr >= len(s.Data) {
		return false
	}
	s.Data[s.Ptr] = dat
	s.Ptr += 1
	return true
}

func (s *stack[T]) Pop(def T) (T, bool) {
	if s.Ptr <= 0 {
		return def, false
	}
	s.Ptr -= 1
	return s.Data[s.Ptr], true
}

func (s *stack[T]) Len() int {
	return s.Ptr
}
