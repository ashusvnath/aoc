package main

import (
	"log"
	"math"
	"math/cmplx"
	"regexp"
	"sort"
)

type PathFinder struct {
	visited Set[complex128]
	path    []complex128
	g       *Grid
	//knownDistances   map[complex128]int
	backtrackedNodes Set[complex128]
}

func (pf *PathFinder) FindPath(sLen int) {
	g := pf.g
	currentIdx := g.start
	currentHeight := 0
	pf.visited.Add(currentIdx)
	for currentIdx != g.end {
		prevIdx := currentIdx
		neighbours := pf.rankedNeighbours(currentIdx)
		log.Printf("idx: %v, height: %v, neighbours, :%#v", currentIdx, currentHeight, neighbours)
		for _, n := range neighbours {
			if !pf.visited.Contains(n) {
				currentHeight = pf.g.mat[n]
				pf.visited.Add(n)
				currentIdx = n
				pf.path = append(pf.path, n)
				break
			}
		}
		if currentIdx == prevIdx {
			//Backtracking because no progress can be made
			backtrackCount := 1
			if pf.backtrackedNodes.Contains(currentIdx) {
				log.Printf("Backtrack: Could be in a loop. Increasing backtrack step to %d.", backtrackCount)
				backtrackCount++
			}
			l := len(pf.path)
			log.Printf("Backtrack: No progress when checking for ranked neighbours of %v. Backtracking by 1 step.", currentIdx)
			//delete(pf.visited, currentIdx)
			pf.backtrackedNodes.Add(currentIdx)
			currentIdx = pf.path[l-1-backtrackCount]
			newPath := make([]complex128, len(pf.path)-1-backtrackCount)
			copy(newPath, pf.path)
			newPath = append(newPath, currentIdx)
			pf.path = newPath
		}
	}
	pf.Shorten(sLen)
}

func (pf *PathFinder) VisualizePath() string {
	g := pf.g
	path := make([]complex128, len(pf.path))
	copy(path, pf.path)
	path = append(path, g.end)
	displayData := []rune(string(pf.g.rawData))
	displayByteMap := map[complex128]rune{
		-1: '\u2190', 1: '\u2192',
		-1i: '\u2191', 1i: '\u2193',
		0: 'E',
	}
	for x, idx := range pf.path {
		nextEntry := path[x+1]
		d := displayByteMap[nextEntry-idx]
		if d == 0 {
			d = '?'
		}
		displayIdx := int(imag(idx))*(pf.g.cols+1) + int(real(idx))
		displayData[displayIdx] = d
	}
	return regexp.MustCompile(`[bd-z]`).ReplaceAllString(string(displayData), "·")
}

func (pf *PathFinder) dist(idx1, idx2 complex128) int {
	dx := math.Abs(real(idx1 - idx2))
	dy := math.Abs(imag(idx1 - idx2))
	v := int(dx + dy)
	return v
}

func (pf *PathFinder) rankedNeighbours(idx complex128) []complex128 {
	g := pf.g
	currentHeight := g.mat[idx]
	_n := g.Neighbours(idx)
	log.Printf("Neighbours: of %v Unranked: %#v", idx, _n)
	n := pf.filter(_n, currentHeight)
	sort.Slice(n, func(i, j int) bool {
		lHeight := pf.g.mat[n[i]]
		rHeight := pf.g.mat[n[j]]
		//Rank by height otherwise, taller better
		if lHeight > rHeight {
			return true
		} else if lHeight < rHeight {
			return false
		} else {
			return pf.dist(n[i], g.end) < pf.dist(n[j], g.end)
		}
	})
	return n
}

func (pf *PathFinder) filter(in []complex128, h int) []complex128 {
	result := []complex128{}
	for _, idx := range in {
		nHeight := pf.g.mat[idx]
		if nHeight-h > 1 || pf.visited.Contains(idx) {
			continue
		}
		result = append(result, idx)
	}
	log.Printf("Neighbours: Filtered: %#v", result)
	return result
}

func NewPathFinder(g *Grid) *PathFinder {
	return &PathFinder{
		visited: make(Set[complex128]),
		path:    []complex128{},
		g:       g,
		//knownDistances:   make(map[int]int),
		backtrackedNodes: make(Set[complex128]),
		//queue:            NewQueue[int](),
	}
}

func (pf *PathFinder) Shorten(shorteningRangeMax int) {
	log.Println("Shortening: Starting procedure")
	shortened := true
	shorteningRange := 3
	for shorteningRange < shorteningRangeMax {
		log.Printf("Shortening: looking for segments of length %d", shorteningRange+1)
		shortened = false
		for i := shorteningRange; i < len(pf.path); i++ {
			idx1 := pf.path[i-shorteningRange]
			idx2 := pf.path[i]
			h1, h2 := pf.g.mat[idx1], pf.g.mat[idx2]
			hDiff := h1 - h2
			diff := idx1 - idx2
			diff = diff * diff
			if (real(diff) == 0 || imag(diff) == 0) && hDiff == 0 || hDiff == -1 {
				log.Printf("Shortening: Trying to shorten between %v(%v):%v, %v(%v):%v. ", idx1, i-shorteningRange, h1, idx2, i, h2)
				oldPathSegment := pf.path[i-shorteningRange : i+1]
				newPathSegment, shortened := pf.createPath(idx1, idx2)
				if shortened && len(oldPathSegment) > len(newPathSegment) {
					// Before shortening
					// 0.....i-sR....i......
					// After shortening
					// 0.....i-sR-1,newSeg,i+1........
					log.Printf("Shortening: old path:%#v ", oldPathSegment)
					log.Printf("Shortening: new Path: %#v", newPathSegment)
					newPath := append([]complex128{}, pf.path[:i-shorteningRange]...)
					newPath = append(newPath, newPathSegment...)
					newPath = append(newPath, pf.path[i+1:]...)
					pf.path = newPath
					log.Printf("Shortening: Restarting loop to shorten segments of size %d", shorteningRange)
					break
				}
			}
		}
		if !shortened {
			log.Printf("Shortening: Increasing search length!!")
			shorteningRange++
		}
	}
}

func (pf *PathFinder) createPath(idx1, idx2 complex128) ([]complex128, bool) {
	connection := []complex128{}
	h := pf.g.mat[idx1]
	delta := idx2 - idx1
	start, end := idx1, idx2
	d, _ := cmplx.Polar(delta)
	delta = complex(real(delta)/d, imag(delta)/d)
	i := 0i
	log.Printf("Shortening: ]]]]]]]]]]]] : For new path start:%v, end:%v, delta:%v", start, end, delta)

	for i = start; i != end && pf.g.mat[i] == h; i += delta {
		connection = append(connection, i)
	}
	if i != end {
		log.Println("Shortening: ]]]]]]]]]]]] : Direct path with same height not found")
		return nil, false
	}
	connection = append(connection, end)
	return connection, len(connection) > 0
}
