package main

import "log"

type Rope struct {
	allKnots map[int]*Knot
}

func (r *Rope) Move(direction string, count int) {
	log.Printf("Rope: Moving: %v, %v", direction, count)
	r.allKnots[1].Move(direction, count)
}

func (r *Rope) RegisterRecorderByKnotIdx(recorder Notifiable[*Knot], knotNumber int) {
	knotToFollow := r.allKnots[knotNumber]
	knotToFollow.Register(recorder)
}

func NewRope(numKnots int) *Rope {
	allKnots := make(map[int]*Knot)
	allKnots[1] = NewKnot(1)
	for i := 2; i <= numKnots; i++ {
		allKnots[i] = NewKnot(i)
		allKnots[i-1].Register(allKnots[i].Follow)

	}
	rope := &Rope{allKnots}
	return rope
}
