package models

import (
	"day12/utility"
	"strings"
)

type Grid struct {
	rawData         []byte
	mat             map[complex128]int
	idxsByHeight    map[int]*utility.Set[complex128]
	start, end      complex128
	rows, cols      int
	knownNeighbours map[complex128][]complex128
}

func (g *Grid) End() complex128 {
	return g.end
}

func (g *Grid) Start() complex128 {
	return g.start
}

func (g *Grid) Columns() int {
	return g.cols
}

func (g *Grid) RawData() []byte {
	return g.rawData
}

func (g *Grid) HeightAt(location complex128) int {
	h, found := g.mat[location]
	if found {
		return h
	}
	return -1
}

func (g *Grid) GetLocationsAtHeight(height int) *utility.Set[complex128] {
	return g.idxsByHeight[height]
}

func (g *Grid) Neighbours(in complex128) []complex128 {
	neighbours := []complex128{in + 1, in + 1i, in - 1, in - 1i}
	result := g.knownNeighbours[in]
	if result != nil {
		return result
	}

	result = []complex128{}
	for _, gp := range neighbours {
		c, r := int(real(gp)), int(imag(gp))
		if r < 0 || c < 0 || r >= g.rows || c >= g.cols {
			continue
		}
		result = append(result, gp)
	}
	g.knownNeighbours[in] = result
	return result
}

func ParseGrid(data []byte) *Grid {
	input := strings.TrimSuffix(string(data), "\n")
	start, end := -1i, -1i
	idxsByHeight := make(map[int]*utility.Set[complex128])
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
			idxsByHeight[height] = utility.NewSet[complex128]()
		}
		idxsByHeight[height].Add(gp)
		gp += 1
	}
	pi := &Grid{
		rawData:         data,
		mat:             matrix,
		idxsByHeight:    idxsByHeight,
		start:           start,
		end:             end,
		rows:            linesSeen + 1,
		cols:            int(real(gp)),
		knownNeighbours: make(map[complex128][]complex128),
	}
	return pi
}
