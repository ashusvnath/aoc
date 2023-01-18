package main

import "log"

var moveDelta = map[string]complex128{
	"R": 1,
	"L": -1,
	"U": 1i,
	"D": -1i,
}

type Rope struct {
	head          *Knot
	tailPositions Set[complex128]
}

func (r *Rope) Move(direction string, count int) {
	r.head.Move(direction, count)
}

func (r *Rope) Record(tailKnot *Knot) {
	log.Printf("Tail moved to : %v", tailKnot.position)
	r.tailPositions.Add(tailKnot.position)
}

func NewRope(numKnots int) *Rope {
	positions := make(Set[complex128])
	positions.Add(0)
	tail := NewKnot()
	prev := tail
	var head *Knot
	for i := 1; i < numKnots; i++ {
		head = NewKnot()
		head.Register(prev.Follow)
		prev = head
	}
	rope := &Rope{head, positions}
	tail.Register(rope.Record)
	return rope
}
