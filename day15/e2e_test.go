package main

import (
	"day14/assert"
	"day15/models"
	"day15/parser"
	"testing"
)

func TestE2E(t *testing.T) {
	//t.Skip()
	input := string(readFile("test.txt"))
	part1 := models.NewSensorProximityObserver(10)
	part2 := models.NewMissingBeaconLocator(20)

	grid := parser.Parse(input, part1.Listen, part2.Process)
	assert.False(grid == nil, t)
	assert.Equal(26, part1.Count(), t)
	//assert.Equal(1, len(tfs), t)
}
