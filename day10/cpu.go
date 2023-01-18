package main

import (
	"regexp"
	"strconv"
)

type CPURecord interface {
	Cycle() int
	RegisterX() int
}

var addInstRegex *regexp.Regexp

func init() {
	addInstRegex = regexp.MustCompile(`addx ([-]?\d+)`)
}

type CPU struct {
	x     int
	clock *Clock
	o     *Observable[CPURecord]
}

func (c *CPU) Tick() {
	c.clock.Tick()
	c.o.Notify()
}

func (c *CPU) Cycle() int {
	return c.clock.cycle
}

func (c *CPU) RegisterX() int {
	return c.x
}

func (c *CPU) Execute(instructions []string) {
	for _, instruction := range instructions {
		switch {
		case addInstRegex.Match([]byte(instruction)):
			c.Tick()
			c.Tick()
			submatches := addInstRegex.FindSubmatch([]byte(instruction))
			op, _ := strconv.Atoi(string(submatches[1]))
			c.x += op

		case instruction == "noop":
			c.Tick()
		}
	}
}

func (c *CPU) Register(r Recorder) {
	c.o.Register(r.Record)
}

func NewCPU(clk *Clock) *CPU {
	cpu := &CPU{1, clk, nil}
	obs := &Observable[CPURecord]{cpu, nil}
	cpu.o = obs
	return cpu
}
