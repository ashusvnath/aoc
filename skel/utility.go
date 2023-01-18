package main

type Set[T comparable] map[T]bool

func (s *Set[T]) Contains(val T) bool {
	_, ok := (*s)[val]
	return ok
}

func (s *Set[T]) Add(elem T) {
	(*s)[elem] = true
}

func (s *Set[T]) AsSlice() []T {
	result := make([]T, len(*s))
	idx := 0
	for k := range *s {
		result[idx] = k
		idx++
	}
	return result
}

type Notifiable[T any] func(T)

type Observable[T any] struct {
	observed  T
	observers []Notifiable[T]
}

func (o *Observable[T]) Register(n Notifiable[T]) {
	o.observers = append(o.observers, n)
}

func (o *Observable[T]) Notify() {
	for _, notify := range o.observers {
		notify(o.observed)
	}
}
