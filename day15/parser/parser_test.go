package parser

import (
	"day15/assert"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("Parse co-ordinates of sensor and beacon", func(t *testing.T) {
		input := "Sensor at x=2, y=18: closest beacon is at x=-2, y=15"

		sensor, beacon := ParseLine(input)

		assert.Equal(2+18i, sensor, t)
		assert.Equal(-2+15i, beacon, t)
	})
}
