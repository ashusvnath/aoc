package utility

type BFSNeighbourFunction[T any] func(T) []T
type Pair[T comparable] [2]T

type BFS[T comparable] struct {
	visited    Set[T]
	queue      Queue[[]T]
	knownPaths map[Pair[T]][]T
}

func (b *BFS[T]) FindShortestPath(start, end T, neighboursFunc BFSNeighbourFunction[T]) []T {
	keyPair := Pair[T]{start, end}
	if value, found := b.knownPaths[keyPair]; found {
		return value
	}
	path := []T{}

	b.queue.Enqueue([]T{start})

	for b.queue.Len() > 0 {
		path := b.queue.Dequeue()
		currentIdx := path[len(path)-1]
		neighbours := neighboursFunc(currentIdx)
		for _, n := range neighbours {
			if n == end {
				path = append(path, n)
				b.knownPaths[keyPair] = path
				b.Clear()
				return path
			}
			if !b.visited.Contains(n) {
				nPair := Pair[T]{n, end}
				if p, found := b.knownPaths[nPair]; found {
					path = append(path, p...)
					b.knownPaths[keyPair] = path
					b.Clear()
					return path
				} else {
					b.visited.Add(n)
					newPath := append([]T{}, path...)
					newPath = append(newPath, n)
					b.queue.Enqueue(newPath)
				}
			}
		}
	}
	return path
}

func (b *BFS[T]) Clear() {
	b.visited = NewMapSet[T]()
	b.queue.Clear()
}

func NewBFS[T comparable]() *BFS[T] {
	return &BFS[T]{
		queue:      NewQueue[[]T](),
		visited:    NewMapSet[T](),
		knownPaths: make(map[Pair[T]][]T),
	}
}
