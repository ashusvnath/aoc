package utility

import "log"

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
