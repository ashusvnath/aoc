package utility

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
