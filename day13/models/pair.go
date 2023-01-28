package models

type Pair struct {
	left, right List
}

func (p *Pair) Equal(other *Pair) bool {
	leftSame := p.left.Compare(other.left) == Equal
	rightSame := p.right.Compare(other.right) == Equal
	return leftSame && rightSame
}

func (p *Pair) IsOrderedCorrectly() bool {
	return p.left.Compare(p.right) == Less
}

func NewPair(left, right List) *Pair {
	return &Pair{left, right}
}
