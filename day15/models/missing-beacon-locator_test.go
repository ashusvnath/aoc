package models

import (
	"day15/assert"
	"testing"
)

func TestMissingBeaconLocator(t *testing.T) {
	t.Run("NewMissingBeaconLocator should create a missing beacon locator", func(t *testing.T) {
		mbl := NewMissingBeaconLocator(20)

		assert.True(mbl != nil, t)
		assert.Equal(20, mbl.size, t)
	})

	t.Run("Process should add sensor and beacon information to its scanlines", func(t *testing.T) {
		sLoc := Location(complex(10, 10))
		bLoc := Location(complex(10, 15))

		mbl := NewMissingBeaconLocator(20)

		mbl.Process(NewSensor(sLoc, bLoc))
		assert.Equal(11, len(mbl.scanLines), t)
		for i := 0.0; i < 5.0; i++ {
			assert.Equal(window{5 + i, 15 - i}, mbl.scanLines[10-i].exclusions[0], t)
		}
	})
}
