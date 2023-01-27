package main

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

func (s *Set[T]) AsSlice() []T {
	result := make([]T, len(s.s))
	idx := 0
	for k := range s.s {
		result[idx] = k
		idx++
	}
	return result
}

func (s *Set[T]) Len() int {
	return len(s.s)
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{make(map[T]bool)}
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

type Queue[T any] struct {
	q []T
}

func (q *Queue[T]) Push(elem T) {
	q.q = append(q.q, elem)
}

func (q *Queue[T]) Pop() T {
	if len(q.q) == 0 {
		return *new(T)
	}
	result := q.q[0]
	q.q = q.q[1:]
	return result
}

func (q *Queue[T]) Len() int {
	return len(q.q)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.q) == 0
}

func (q *Queue[T]) Clear() {
	q.q = []T{}
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func NewQueueN[T any](n int) *Queue[T] {
	return &Queue[T]{make([]T, n)}
}
