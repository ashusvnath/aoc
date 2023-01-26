package main

import "testing"

func TestParse(t *testing.T) {
	input := []byte(`aSbcde
fghEjk
lmnopq`)
	t.Run("Parse input", func(t *testing.T) {
		pi := Parse(input)
		t.Run("Should parse start cell correctly", func(t *testing.T) {
			assertEqual(1, pi.start, t)
		})
		t.Run("Should parse end cell correctly", func(t *testing.T) {
			assertEqual(9, pi.end, t)
		})
		t.Run("Should parse rows correctly", func(t *testing.T) {
			assertEqual(3, pi.rows, t)
		})
		t.Run("Should parse cols correctly", func(t *testing.T) {
			assertEqual(6, pi.cols, t)
		})
		t.Run("Should parse each cell correctly", func(t *testing.T) {
			idx := 0
			for _, c := range input {
				if c == '\n' {
					continue
				}
				if c == 'S' {
					c = 'a'
				}
				if c == 'E' {
					c = 'z'
				}
				assertEqual(int(c)-'a', pi.GetByIdx(idx), t)
				idx++
			}

		})
	})

	t.Run("Parsed input", func(t *testing.T) {
		pi := Parse(input)
		t.Run("should get element by index", func(t *testing.T) {
			idx := 8
			assertEqual(7, pi.GetByIdx(idx), t)
		})

		t.Run("should get element by row and column", func(t *testing.T) {
			row, col := 2, 4
			assertEqual(15, pi.GetByRC(row, col), t)
		})

		t.Run("should produce neighbours", func(t *testing.T) {
			t.Run("for idx 0, top left", func(t *testing.T) {
				n := pi.Neighbours(0)
				assertEqual(2, len(n), t)
				assertEqual(6, n[0], t)
				assertEqual(1, n[1], t)
			})

			t.Run("for idx 3, top mid", func(t *testing.T) {
				n := pi.Neighbours(3)
				assertEqual(3, len(n), t)
				assertEqual(9, n[0], t)
				assertEqual(4, n[1], t)
				assertEqual(2, n[2], t)
			})

			t.Run("for idx 5, top right", func(t *testing.T) {
				n := pi.Neighbours(5)
				assertEqual(2, len(n), t)
				assertEqual(4, n[1], t)
				assertEqual(11, n[0], t)
			})

			t.Run("for idx 6, left mid", func(t *testing.T) {
				n := pi.Neighbours(6)

				assertEqual(3, len(n), t)
				assertEqual(12, n[0], t)
				assertEqual(0, n[1], t)
				assertEqual(7, n[2], t)
			})

			t.Run("for idx 11, right mid", func(t *testing.T) {
				n := pi.Neighbours(11)

				assertEqual(3, len(n), t)
				assertEqual(17, n[0], t)
				assertEqual(5, n[1], t)
				assertEqual(10, n[2], t)

			})

			t.Run("for idx 12, bottom left", func(t *testing.T) {
				n := pi.Neighbours(12)
				assertEqual(2, len(n), t)
				assertEqual(6, n[0], t)
				assertEqual(13, n[1], t)
			})

			t.Run("for idx 15, bottom mid", func(t *testing.T) {
				n := pi.Neighbours(15)

				assertEqual(3, len(n), t)
				assertEqual(9, n[0], t)
				assertEqual(16, n[1], t)
				assertEqual(14, n[2], t)
			})

			t.Run("for idx 17, bottom right", func(t *testing.T) {
				n := pi.Neighbours(17)
				assertEqual(2, len(n), t)
				assertEqual(11, n[0], t)
				assertEqual(16, n[1], t)
			})

			t.Run("for id 9, middle",func(t *testing.T) {
				n := pi.Neighbours(9)
				assertEqual(4, len(n), t)
				assertEqual(15, n[0], t)
				assertEqual(3, n[1], t)
				assertEqual(10, n[2], t)
				assertEqual(8, n[3], t)
			})
		})
	})

}
