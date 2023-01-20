package main

import "testing"

func TestParse(t *testing.T) {
	input := []byte(`aSbcde
fghijE`)
	t.Run("Parse input", func(t *testing.T) {
		pi := Parse(input)
		t.Run("Should parse start cell correctly", func(t *testing.T) {
			assertEqual(1, pi.start, t)
		})
		t.Run("Should parse end cell correctly", func(t *testing.T) {
			assertEqual(11, pi.end, t)
		})
		t.Run("Should parse rows correctly", func(t *testing.T) {
			assertEqual(2, pi.rows, t)
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
				assertEqual(int(c)-'a', pi.GetByIdx(idx), t)
				idx++
			}

		})
	})

}
