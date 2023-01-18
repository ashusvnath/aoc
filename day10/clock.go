package main

type Clock struct {
	cycle int
}

func (clk *Clock) Tick() {
	clk.cycle += 1
}

func NewClock() *Clock {
	return &Clock{0}
}
