package main

type Sprite struct {
	pos int
}

func (s *Sprite) Overlaps(x int) bool {
	return s.pos-2 < x && x < s.pos+2
}
