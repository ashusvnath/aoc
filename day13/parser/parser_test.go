package parser

import (
	"day13/assert"
	. "day13/models"
	"runtime"
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("Should parse empty list", func(t *testing.T) {
		input := `[]`

		actual := Parse(input)
		emptyList := List([]Entry{})
		assertEqualEntry(emptyList, actual, t)
	})

	t.Run("Should parse list with one integer value", func(t *testing.T) {
		input := `[1]`

		actual := Parse(input)
		expected := List([]Entry{Int(1)})
		assertEqualEntry(expected, actual, t)
	})

	t.Run("Should parse list with multiple integer values", func(t *testing.T) {
		input := `[[10,2,333,4]]`

		actual := Parse(input)
		expected := List{List{Int(10), Int(2), Int(333), Int(4)}}
		assertEqualEntry(expected, actual, t)
	})

	t.Run("Should parse list with integers and an embedded list", func(t *testing.T) {
		input := `[1,2,[3],4]`

		actual := Parse(input)
		expected := List([]Entry{Int(1), Int(2), List{Int(3)}, Int(4)})
		assertEqualEntry(expected, actual, t)
	})

	t.Run("Should parse nested list", func(t *testing.T) {
		input := `[[[]]]`

		actual := Parse(input)
		expected := List{List{List{}}}
		assertEqualEntry(expected, actual, t)
	})
}

func TestParsePairs(t *testing.T) {
	t.Run("Should parse pairs", func(t *testing.T) {
		input := "[1,2,3,[[]],4]\n[4,[],[],5,6]\n\n[]\n[9]"
		expectedLeft := List{Int(1), Int(2), Int(3), List{List{}}, Int(4)}
		expectedRight := List{Int(4), List{}, List{}, Int(5), Int(6)}
		expected0 := NewPair(expectedLeft, expectedRight)
		expected1 := NewPair(List{}, List{Int(9)})
		actual := ParsePairs(input)

		assert.Equal(2, len(actual), t)
		assert.True(expected0.Equal(actual[0]), t)
		assert.True(expected1.Equal(actual[1]), t)
	})
}

func assertEqualEntry(expected, actual Entry, t *testing.T) {
	_, file, line, _ := runtime.Caller(1)
	if expected.Compare(actual) != Equal {
		t.Errorf("\n%s:%d\nExpected : %v\nActual   : %v\n", file, line, expected, actual)
	}
}
