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
	obs := models.NewSensorProximityObserver(10)

	grid := parser.Parse(input, obs.Listen)
	assert.False(grid == nil, t)
	assert.Equal(26, obs.Count(), t)
}
