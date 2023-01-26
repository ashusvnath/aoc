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

func NewSet[T comparable]() Set[T] {
	return make(map[T]bool)
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

type Queue[T any] []T

func (q *Queue[T]) Push(elem T) {
	typedQ := Queue[T](append([]T(*q), elem))
	q = &typedQ
}

func (q *Queue[T]) Pop() T {
	if len(*q) == 0 {
		return *new(T)
	}
	typedQ := []T(*q)
	result := typedQ[0]
	typedQ = typedQ[1:]
	newQ := Queue[T](typedQ)
	q = &newQ
	return result
}

func (q *Queue[T]) Len() int {
	return len(*q)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

func NewQueue[T any]() *Queue[T] {
	newQ := make([]T, 0)
	q := Queue[T](newQ)
	return &q
}

func NewQueueN[T any](n int) *Queue[T] {
	newQ := make([]T, n)
	q := Queue[T](newQ)
	return &q
}
