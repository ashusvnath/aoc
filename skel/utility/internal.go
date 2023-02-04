package utility

type linearNode[T any] struct {
	data T
	next *linearNode[T]
}
