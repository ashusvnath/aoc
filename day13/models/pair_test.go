package models

import (
	"day13/assert"
	"testing"
)

func TestPair(t *testing.T) {
	t.Run("IsOrderedCorrectly", func(t *testing.T) {
		t.Run("should return true when left < right", func(t *testing.T) {
			p := NewPair(List{List{}, List{}}, List{Int(1)})
			assert.True(p.IsOrderedCorrectly(), t)
		})

		t.Run("should return false left >= right", func(t *testing.T) {
			p := NewPair(List{Int(1), List{Int(2)}, List{List{}}}, List{Int(1), List{List{List{}}}})
			assert.False(p.IsOrderedCorrectly(), t)
		})
	})

}
