package main

import "strings"

type ParsedInput struct {
	rawData      []byte
	mat          [][]int
	idxsByHeight map[int][]int
	start, end   int
	rows, cols   int
}

func (pi *ParsedInput) GetByIdx(idx int) int {
	row := idx / pi.cols
	col := idx % pi.cols
	return pi.mat[row][col]
}

func (pi *ParsedInput) GetByRC(r, c int) int {
	return pi.mat[r][c]
}

func (pi *ParsedInput) Neighbours(idx int) []int {
	row := idx / pi.cols
	col := idx % pi.cols
	neighbours := [][2]int{{row + 1, col}, {row - 1, col}, {row, col + 1}, {row, col - 1}}
	result := []int{}
	for _, rc := range neighbours {
		row, col := rc[0], rc[1]
		if row < 0 || col < 0 || row >= pi.rows || col >= pi.cols {
			continue
		}
		result = append(result, row*pi.rows+col)
	}
	return result
}

func Parse(data []byte) *ParsedInput {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	l := len(lines)
	matrix := make([][]int, l)
	start, end := -1, -1
	idxsByHeight := make(map[int][]int)
	for row, lineString := range lines {
		matrix[row] = make([]int, len(lineString))
		for col, c := range lineString {
			height := int(c) - 'a'
			matrix[row][col] = height
			idx := row*(len(lineString)) + col
			if height == -14 {
				start = idx
			}
			if height == -28 {
				end = idx
			}
			idxsByHeight[height] = append(idxsByHeight[height], idx)
		}
	}
	pi := &ParsedInput{
		rawData:      data,
		mat:          matrix,
		idxsByHeight: idxsByHeight,
		start:        start,
		end:          end,
		rows:         len(matrix),
		cols:         len(matrix[0]),
	}
	return pi
}
