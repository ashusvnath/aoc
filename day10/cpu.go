package main

import (
	"log"
	"regexp"
	"strconv"
)

var addInstRegex *regexp.Regexp

func init() {
	addInstRegex = regexp.MustCompile(`addx ([-]?\d+)`)
}

type CPU struct {
	X              int
	recordedValues map[uint]int
	recordAt       Set[uint]
	cycle          uint
}

func (c *CPU) Tick() {
	c.cycle += 1
	if c.recordAt.Contains(c.cycle) {
		c.recordedValues[c.cycle] = c.X
	}
}

func (c *CPU) Execute(instructions []string) {
	for _, instruction := range instructions {
		switch {
		case addInstRegex.Match([]byte(instruction)):
			c.Tick()
			c.Tick()
			submatches := addInstRegex.FindSubmatch([]byte(instruction))
			op, _ := strconv.Atoi(string(submatches[1]))
			c.X += op

		case instruction == "noop":
			c.Tick()
		}
	}
}

func (c *CPU) RecordRegsiterValueAtCycle(cycle ...uint) {
	for _, cycleIdx := range cycle {
		c.recordAt.Add(cycleIdx)
	}
}

func NewCPU() *CPU {
	return &CPU{1, make(map[uint]int), make(Set[uint]), 0}
}

func Part1(cpu *CPU) int {
	result := 0
	for cycle, value := range cpu.recordedValues {
		diff := int(cycle) * value
		result += diff
		log.Printf("cycle: %3d , value: %3d, diff: %10v, cumulative:%10v", cycle, value, diff, result)
	}
	return result
}
