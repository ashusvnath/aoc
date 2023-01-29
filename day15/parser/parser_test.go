package parser

import (
	"day15/assert"
	"day15/models"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("Parse co-ordinates of sensor and beacon", func(t *testing.T) {
		input := "Sensor at x=2, y=18: closest beacon is at x=-2, y=15"

		sensor, beacon := ParseLine(input)

		assert.Equal(2+18i, sensor, t)
		assert.Equal(-2+15i, beacon, t)
	})

	t.Run("", func(t *testing.T) {
		input := `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16`

		grid := Parse(input, func(*models.Sensor) {})

		assert.Equal(models.SENSOR, grid.ObjectAt(models.Location(2+18i)).Type(), t)
		assert.Equal(models.SENSOR, grid.ObjectAt(models.Location(2+18i)).Type(), t)
		assert.Equal(models.SENSOR, grid.ObjectAt(models.Location(2+18i)).Type(), t)
		assert.Equal(models.SENSOR, grid.ObjectAt(models.Location(2+18i)).Type(), t)
	})
}
