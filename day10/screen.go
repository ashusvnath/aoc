package main

import "bytes"

type Screen struct {
	out        *bytes.Buffer
	on, off    byte
	rows, cols int
}

func (sc *Screen) Report() string {
	return sc.out.String()
}

func (sc *Screen) Record(r CPURecord) {
	idx := r.Cycle() - 1
	register := r.RegisterX()
	spr := Sprite{register}
	if spr.Overlaps(idx % sc.cols) {
		sc.out.WriteByte(sc.on)
	} else {
		sc.out.WriteByte(sc.off)
	}
	if (idx+1)%sc.cols == 0 {
		sc.out.WriteByte('\n')
	}
}

func NewScreen(rows, cols int, on, off byte) *Screen {
	buf := make([]byte, rows*(cols+1))
	return &Screen{bytes.NewBuffer(buf), on, off, rows, cols}
}
