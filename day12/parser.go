package main

import "strings"

type Grid struct {
	rawData      []byte
	mat          map[complex128]int
	idxsByHeight map[int]*Set[complex128]
	start, end   complex128
	rows, cols   int
}

func (g *Grid) Neighbours(in complex128) []complex128 {
	neighbours := []complex128{in + 1i, in - 1i, in + 1, in - 1}
	result := []complex128{}
	for _, gp := range neighbours {
		c, r := int(real(gp)), int(imag(gp))
		if r < 0 || c < 0 || r >= g.rows || c >= g.cols {
			continue
		}
		result = append(result, gp)
	}
	return result
}

func Parse(data []byte) *Grid {
	input := strings.TrimSuffix(string(data), "\n")
	start, end := -1i, -1i
	idxsByHeight := make(map[int]*Set[complex128])
	matrix := make(map[complex128]int)
	linesSeen := 0
	gp := 0i
	for _, c := range input {
		height := int(c) - 'a'
		if c == 'S' {
			start = gp
			height = 0
		}
		if c == 'E' {
			end = gp
			height = 25
		}
		if c == '\n' {
			gp = complex(0, imag(gp)+1)
			linesSeen += 1
			continue
		}

		matrix[gp] = height
		if idxsByHeight[height] == nil {
			s := make(Set[complex128])
			idxsByHeight[height] = &s
		}
		idxsByHeight[height].Add(gp)
		gp += 1
	}
	pi := &Grid{
		rawData:      data,
		mat:          matrix,
		idxsByHeight: idxsByHeight,
		start:        start,
		end:          end,
		rows:         linesSeen + 1,
		cols:         int(real(gp)),
	}
	return pi
}
