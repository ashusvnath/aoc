package main

import (
	"math"
	"math/cmplx"
)

type Knot struct {
	position  complex128
	followers []Notifiable[*Knot]
}

func (k *Knot) Move(direction string, count int) {
	for i := 0; i < count; i++ {
		k.position += moveDelta[direction]
		k.Notify()
	}
}

func (k *Knot) Register(n Notifiable[*Knot]) {
	k.followers = append(k.followers, n)
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
		return
	}
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

func NewKnot() *Knot {
	return &Knot{0, nil}
}
