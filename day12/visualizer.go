package main

import (
	"log"
	"regexp"
)

func VisualizePath(g *Grid, inPath []complex128, start complex128) string {
	path := make([]complex128, len(inPath))
	copy(path, inPath)
	path = append(path, g.end)
	base := regexp.MustCompile(`[a-zS]`).ReplaceAllString(string(g.rawData), "\u00B7")
	displayData := []rune(base)
	displayByteMap := map[int]map[complex128]rune{
		-2: {
			-1: 'a', 1: 'd',
			1i: 's', -1i: 'w',
		},
		-1: {
			-1: '⇠', 1: '⇢',
			1i: '⇣', -1i: '͎',
		},

		1: {
			-1: '<', 1: '>',
			-1i: '^', 1i: 'V',
			0: 'E',
		},

		0: {
			-1: '←', 1: '→',
			1i: '↓', -1i: '↑',
			0: 'S',
		},
	}
	prev := start
	for x, idx := range inPath[:len(inPath)-1] {
		nextEntry := path[x+1]
		hDiff := g.mat[idx] - g.mat[prev]
		d := displayByteMap[hDiff][idx-prev]
		if d == 0 {
			d = '?'
			log.Printf("idx: %v, nextEntry: %v, prev: %v, hDiff: %d", idx, nextEntry, prev, hDiff)
		}
		displayIdx := int(imag(idx))*(g.cols+1) + int(real(idx))
		displayData[displayIdx] = d
		prev = idx
	}

	displayIdx := int(imag(start))*(g.cols+1) + int(real(start))
	displayData[displayIdx] = 'S'
	return string(displayData)
}
