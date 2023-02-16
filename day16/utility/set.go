package utility

type Set[T comparable] struct {
	s map[T]bool
}

func (s *Set[T]) Contains(val T) bool {
	_, ok := (s.s)[val]
	return ok
}

func (s *Set[T]) Add(elem T) {
	s.s[elem] = true
}

func (s *Set[T]) Remove(elem T) {
	delete(s.s, elem)
}

func (s *Set[T]) AsSlice() []T {
	result := make([]T, len(s.s))
	idx := 0
	for k := range s.s {
		result[idx] = k
		idx++
	}
	return result
}

func (s *Set[T]) AddAll(in ...T) {
	for _, k := range in {
		s.s[k] = true
	}
}

func (s *Set[T]) Len() int {
	return len(s.s)
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{make(map[T]bool)}
}
