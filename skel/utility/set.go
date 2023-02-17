package utility

const MaxInt = 1<<64 - 1

type Set[T comparable] interface {
	Add(elem T)
	Remove(elem T)
	Contains(elem T) bool
	AsSlice() []T
	AddAll(t ...T)
	Clone() Set[T]
}

type MapSet[T comparable] struct {
	s map[T]bool
}

func (s *MapSet[T]) Contains(val T) bool {
	_, ok := (s.s)[val]
	return ok
}

func (s *MapSet[T]) Add(elem T) {
	s.s[elem] = true
}

func (s *MapSet[T]) AddAll(elems ...T) {
	for _, elem := range elems {
		s.s[elem] = true
	}
}

func (s *MapSet[T]) Remove(elem T) {
	delete(s.s, elem)
}

func (s *MapSet[T]) AsSlice() []T {
	result := make([]T, len(s.s))
	idx := 0
	for k := range s.s {
		result[idx] = k
		idx++
	}
	return result
}

func (s *MapSet[T]) Len() int {
	return len(s.s)
}

func (s *MapSet[T]) Clone() Set[T] {
	copiedMap := make(map[T]bool)
	for elem := range s.s {
		copiedMap[elem] = true
	}
	return &MapSet[T]{copiedMap}
}

func NewMapSet[T comparable]() Set[T] {
	return &MapSet[T]{make(map[T]bool)}
}

type BitMask[T comparable] struct {
	possibilities map[T]uint64
}

type BitMaskSet[T comparable] struct {
	bitmask uint64
	options *BitMask[T]
}

func NewBitmask[T comparable](in []T) *BitMask[T] {
	possibilities := make(map[T]uint64)
	for idx, p := range in {
		possibilities[p] = 1 << idx
	}
	return &BitMask[T]{
		possibilities: possibilities,
	}
}

func NewBitMaskSet[T comparable](bitMask *BitMask[T]) *BitMaskSet[T] {
	return &BitMaskSet[T]{
		bitmask: 0,
		options: bitMask,
	}
}

func (bms *BitMaskSet[T]) Add(elem T) {
	bms.bitmask |= bms.options.possibilities[elem]
}

func (bms *BitMaskSet[T]) AddAll(elems ...T) {
	for _, elem := range elems {
		bms.bitmask |= bms.options.possibilities[elem]
	}
}

func (bms *BitMaskSet[T]) Contains(elem T) bool {
	return bms.bitmask&bms.options.possibilities[elem] != 0
}

func (bms *BitMaskSet[T]) Remove(elem T) {
	bms.bitmask &= (MaxInt - bms.options.possibilities[elem])
}

func (bms *BitMaskSet[T]) AsSlice() []T {
	var result []T
	for k := range bms.options.possibilities {
		if bms.bitmask&bms.options.possibilities[k] > 0 {
			result = append(result, k)
		}
	}
	return result
}

func (bms *BitMaskSet[T]) Clone() Set[T] {
	return &BitMaskSet[T]{
		options: bms.options,
		bitmask: bms.bitmask,
	}
}
