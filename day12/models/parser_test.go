package models

import (
	. "day12/utility"
	"testing"
)

func TestParse(t *testing.T) {
	input := []byte(`aSbcde
fghEjk
lmnopq`)
	t.Run("Parse input", func(t *testing.T) {
		pi := ParseGrid(input)
		t.Run("Should parse start cell correctly", func(t *testing.T) {
			AssertEqual(1, pi.start, t)
		})
		t.Run("Should parse end cell correctly", func(t *testing.T) {
			AssertEqual(3+1i, pi.end, t)
		})
		t.Run("Should parse rows correctly", func(t *testing.T) {
			AssertEqual(3, pi.rows, t)
		})
		t.Run("Should parse cols correctly", func(t *testing.T) {
			AssertEqual(6, pi.cols, t)
		})
		t.Run("Should parse each cell correctly", func(t *testing.T) {
			idx := 0i
			for _, c := range input {
				if c == '\n' {
					idx += 1i
					idx -= complex(real(idx), 0)
					continue
				}
				if c == 'S' {
					c = 'a'
				}
				if c == 'E' {
					c = 'z'
				}
				AssertEqual(int(c)-'a', pi.mat[idx], t)
				idx += 1
			}
		})
	})

	t.Run("Parsed input", func(t *testing.T) {
		pi := ParseGrid(input)

		t.Run("should produce neighbours", func(t *testing.T) {
			t.Run("for idx 0, top left", func(t *testing.T) {
				n := pi.Neighbours(0)
				AssertEqual(2, len(n), t)
				AssertEqual(1, n[0], t)
				AssertEqual(1i, n[1], t)
			})

			t.Run("for idx 3, top mid", func(t *testing.T) {
				n := pi.Neighbours(3)
				AssertEqual(3, len(n), t)
				AssertEqual(4, n[0], t)
				AssertEqual(3+1i, n[1], t)
				AssertEqual(2, n[2], t)
			})

			t.Run("for idx 5, top right", func(t *testing.T) {
				n := pi.Neighbours(5)
				AssertEqual(2, len(n), t)
				AssertEqual(4, n[1], t)
				AssertEqual(5+1i, n[0], t)
			})

			t.Run("for idx 6, left mid", func(t *testing.T) {
				n := pi.Neighbours(1i)

				AssertEqual(3, len(n), t)
				AssertEqual(1+1i, n[0], t)
				AssertEqual(2i, n[1], t)
				AssertEqual(0, n[2], t)
			})

			t.Run("for idx 11, right mid", func(t *testing.T) {
				n := pi.Neighbours(5 + 1i)

				AssertEqual(3, len(n), t)
				AssertEqual(5+2i, n[0], t)
				AssertEqual(4+1i, n[1], t)
				AssertEqual(5, n[2], t)

			})

			t.Run("for idx 12, bottom left", func(t *testing.T) {
				n := pi.Neighbours(2i)
				AssertEqual(2, len(n), t)
				AssertEqual(1+2i, n[0], t)
				AssertEqual(1i, n[1], t)
			})

			t.Run("for idx 15, bottom mid", func(t *testing.T) {
				n := pi.Neighbours(2i + 3)

				AssertEqual(3, len(n), t)
				AssertEqual(2i+4, n[0], t)
				AssertEqual(2i+2, n[1], t)
				AssertEqual(1i+3, n[2], t)
			})

			t.Run("for idx 17, bottom right", func(t *testing.T) {
				n := pi.Neighbours(2i + 5)
				AssertEqual(2, len(n), t)
				AssertEqual(2i+4, n[0], t)
				AssertEqual(1i+5, n[1], t)
			})

			t.Run("for id 9, middle", func(t *testing.T) {
				n := pi.Neighbours(1i + 3)
				AssertEqual(4, len(n), t)
				AssertEqual(1i+4, n[0], t)
				AssertEqual(2i+3, n[1], t)
				AssertEqual(1i+2, n[2], t)
				AssertEqual(3, n[3], t)
			})
		})
	})

}
