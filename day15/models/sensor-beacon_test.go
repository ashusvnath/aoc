package models

import (
	"day15/assert"
	"testing"
)

func TestModels(t *testing.T) {
	sensorLoc := Location(complex(2, 18))
	beaconLoc := Location(complex(-2, 15))

	t.Run("location distance should be same both ways", func(t *testing.T) {
		assert.Equal(7, sensorLoc.Distance(beaconLoc), t)
		assert.Equal(7, beaconLoc.Distance(sensorLoc), t)
	})

	t.Run("Should initialize sensor with location and distance to specified beacon", func(t *testing.T) {
		s := NewSensor(sensorLoc, beaconLoc)
		assert.Equal(7, s.closestBeaconDistance, t)
	})
}
