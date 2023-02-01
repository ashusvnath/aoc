package models

import (
	"day15/assert"
	"testing"
)

func TestScanLine(t *testing.T) {
	t.Run("NewScanLine should create window of specified size", func(t *testing.T) {
		sl := NewScanLine(20, 0)

		assert.True(sl != nil, t)
		assert.False(sl.fullyExcluded, t)
		assert.Equal(0, len(sl.exclusions), t)
		assert.Equal(20, sl.max, t)
	})

	t.Run("Exclude should mark as fully excluded if range exhausted", func(t *testing.T) {
		sl := NewScanLine(20, 0)

		sl.Exclude(-1, 21)
		assert.True(sl.fullyExcluded, t)
	})

	t.Run("Exclude should append to exclusions if no overlap found", func(t *testing.T) {
		sl := NewScanLine(20, 0)

		sl.Exclude(2, 4)
		assert.Equal(1, len(sl.exclusions), t)
		sl.Exclude(7, 11)
		assert.Equal(2, len(sl.exclusions), t)
	})

	t.Run("Exclude should merge and simplify exclusions if overlap found", func(t *testing.T) {
		sl := NewScanLine(20, 0)

		sl.Exclude(2, 4)
		assert.Equal(1, len(sl.exclusions), t)
		sl.Exclude(7, 11)
		assert.Equal(2, len(sl.exclusions), t)
		sl.Exclude(3, 6)
		assert.Equal(1, len(sl.exclusions), t)
	})

	t.Run("Exclude should ignore if out of bounds", func(t *testing.T) {
		sl := NewScanLine(20, 0)

		sl.Exclude(-4, -1)
		assert.Equal(0, len(sl.exclusions), t)
		sl.Exclude(21, 100)
		assert.Equal(0, len(sl.exclusions), t)
	})

	t.Run("Exclude should add window with start = stop ", func(t *testing.T) {
		sl := NewScanLine(20, 0)

		sl.Exclude(15, 15)
		assert.Equal(1, len(sl.exclusions), t)
	})
}
