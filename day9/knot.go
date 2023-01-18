package main

import (
	"log"
	"math"
	"math/cmplx"
)

var moveDelta = map[string]complex128{
	"R": 1,
	"L": -1,
	"U": 1i,
	"D": -1i,
}

type Knot struct {
	idx       int
	position  complex128
	followers []Notifiable[*Knot]
}

func (k *Knot) Move(direction string, count int) {
	log.Printf("Knot%d: Moving: %v, %v", k.idx, direction, count)
	for i := 0; i < count; i++ {
		k.position += moveDelta[direction]
		k.Notify()
	}
	log.Printf("Position: %v", k.position)
}

func (k *Knot) Register(follower Notifiable[*Knot]) {
	k.followers = append(k.followers, follower)
}

func (k *Knot) Notify() {
	for _, notify := range k.followers {
		notify(k)
	}
}

func (k *Knot) Follow(leader *Knot) {
	delta := leader.position - k.position
	distMoved, _ := cmplx.Polar(delta)
	if distMoved < 2 {
		log.Printf("Knot%d: Following:Knot%d Skipped for min distance %v, %v: d: %v", k.idx, leader.idx, leader.position, k.position, distMoved)
		return
	}
	log.Printf("Knot%d: Following: %v", k.idx, leader)
	dx, dy := real(delta), imag(delta)
	if dx != 0 {
		dx = dx / math.Abs(dx)
	}
	if dy != 0 {
		dy = dy / math.Abs(dy)
	}
	k.position += complex(dx, dy)
	k.Notify()
}

func NewKnot(idx int) *Knot {
	return &Knot{idx, 0, nil}
}
