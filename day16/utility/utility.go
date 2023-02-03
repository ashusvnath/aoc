package utility

import (
	"log"
	"math/rand"
	"time"
)

var _rng *rand.Rand

func init() {
	_rng = rand.New(rand.NewSource(time.Now().UnixMicro()))
}

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

type sQueue[T any] struct {
	q []T
}

func (q *sQueue[T]) Enqueue(elem T) {
	q.q = append(q.q, elem)
}

func (q *sQueue[T]) Dequeue() T {
	if len(q.q) == 0 {
		return *new(T)
	}
	result := q.q[0]
	q.q = q.q[1:]
	return result
}

func (q *sQueue[T]) Len() int {
	return len(q.q)
}

func (q *sQueue[T]) IsEmpty() bool {
	return len(q.q) == 0
}

func (q *sQueue[T]) Clear() {
	q.q = []T{}
}

func NewSQueue[T any]() *sQueue[T] {
	return &sQueue[T]{}
}

func NewQueueN[T any](n int) *sQueue[T] {
	return &sQueue[T]{make([]T, n)}
}

type linearNode[T any] struct {
	data T
	next *linearNode[T]
}

type lQueue[T any] struct {
	head, tail *linearNode[T]
	len        int
}

func NewlQueue[T any]() *lQueue[T] {
	return &lQueue[T]{nil, nil, 0}
}

func (lq *lQueue[T]) Len() int {
	return lq.len
}

func (lq *lQueue[T]) Dequeue() T {
	var result T
	if lq.len == 0 || lq.head == nil {
		result = *new(T)
	} else {
		data := lq.head.data
		lq.head = lq.head.next
		if lq.head == nil {
			lq.tail = nil
		}
		result = data
	}
	lq.len -= 1
	return result
}

func (lq *lQueue[T]) Enqueue(elem T) {
	newNode := &linearNode[T]{elem, nil}
	lq.len += 1
	if lq.tail == nil {
		lq.tail = newNode
		lq.head = newNode
	} else {
		lq.tail.next = newNode
		lq.tail = newNode
	}
}

func (lq *lQueue[T]) Clear() {
	lq.head = nil
	lq.tail = nil
	lq.len = 0
}

func (lq *lQueue[T]) IsEmpty() bool {
	return lq.len == 0
}

type Queue[T any] interface {
	Enqueue(T)
	Dequeue() T
	Len() int
	IsEmpty() bool
	Clear()
}

type qGen[T any] func() Queue[T]

// NewQueue is an over the top factory to create a new queue backed by a slice or linked list
func NewQueue[T any]() Queue[T] {
	g1 := func() Queue[T] {
		log.Print("Utility: QConstructor: creating lQueue")
		return NewlQueue[T]()
	}
	g2 := func() Queue[T] {
		log.Print("Utility: QConstructor: creating SQueue")
		return NewSQueue[T]()
	}
	generators := []qGen[T]{g1, g2}
	i := _rng.Int() % len(generators)
	return generators[i]()
}

type Stack[T any] struct {
	head  *linearNode[T]
	count int
}

func (s *Stack[T]) Push(elem T) {
	s.count++
	node := &linearNode[T]{elem, nil}
	node.next = s.head
	s.head = node
}

func (s *Stack[T]) Pop() T {
	if s.head == nil {
		return *new(T)
	}
	s.count--
	elem := s.head.data
	s.head = s.head.next
	return elem
}

func (s *Stack[T]) Len() int {
	return s.count
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil, 0}
}
