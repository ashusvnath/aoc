package utility

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
