package main

import (
	"log"
	"math"
	"math/cmplx"
	"math/rand"
	"sort"
)

type PathFinder struct {
	visited          Set[complex128]
	path             []complex128
	g                *Grid
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

func (pf *PathFinder) rankedNeighbours(idx complex128) []complex128 {
	g := pf.g
	currentHeight := g.mat[idx]
	_n := g.Neighbours(idx)
	p := cmplx.Phase(g.end - idx)
	log.Printf("Neighbours: of %v Unranked: %#v", idx, _n)
	n := pf.filter(_n, currentHeight)
	sort.Slice(n, func(i, j int) bool {
		lHeight := pf.g.mat[n[i]]
		rHeight := pf.g.mat[n[j]]
		//Rank by height, taller better
		if lHeight > rHeight {
			return true
		} else if lHeight < rHeight {
			return false
		} else {
			//When same height, closer to end better
			lDiff := g.end - n[i]
			rDiff := g.end - n[j]
			ld, lp := cmplx.Polar(lDiff)
			rd, rp := cmplx.Polar(rDiff)
			if ld < rd {
				return true
			} else if ld > rd {
				return false
			} else {
				dlp := lp - p
				drp := rp - p
				log.Printf("Neighbours: Ranking closeness between: %v(%.4f %+.4f%+.4f), %v(%+.4f %+.4f%+.4f) at %v(%+0.4f) to %v", n[i], ld, p, dlp, n[j], rd, p, drp, idx, p, g.end)
				if math.Abs(dlp) == math.Abs(drp) {
					randomFlip := rand.Float64() > 0.10
					log.Printf("Neighbours: Ranking:Choosing by random flip: %v", randomFlip)
					return randomFlip
				}
				log.Printf("Neighbours: Ranking:Choosing by absolute")
				return math.Abs(dlp) < math.Abs(drp)
			}
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
		backtrackedNodes: make(Set[complex128]),
	}
}

func (pf *PathFinder) Shorten(shorteningRangeMax int) {
	log.Println("Shortening: Starting procedure")
	shorteningRange := 3
	for shorteningRange < shorteningRangeMax {
		log.Printf("Shortening: looking for segments of length %d", shorteningRange+1)
		shortenedAtLeastOnce := false
		for i := shorteningRange; i < len(pf.path); i++ {
			segStart := pf.path[i-shorteningRange]
			segEnd := pf.path[i]
			hStart, hEnd := pf.g.mat[segStart], pf.g.mat[segEnd]
			hDiff := hEnd - hStart
			diff := segEnd - segStart
			if (real(diff) == 0 || imag(diff) == 0) && (hDiff == 0) {
				log.Printf("Shortening: Trying to shorten between %v(%v):%v, %v(%v):%v. ", segStart, i-shorteningRange, hStart, segEnd, i, hEnd)
				oldPathSegment := pf.path[i-shorteningRange : i+1]
				newPathSegment, shortened := pf.createPath(segStart, segEnd, diff, hStart)
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
					shortenedAtLeastOnce = true
					log.Printf("Shortening: Restarting loop to shorten segments of size %d", shorteningRange)
					break
				}
			}
		}
		if !shortenedAtLeastOnce {
			log.Printf("Shortening: Increasing search length!!")
			shorteningRange++
		}
	}
}

func (pf *PathFinder) createPath(start, end, delta complex128, h int) ([]complex128, bool) {
	connection := []complex128{start}
	d, _ := cmplx.Polar(delta)
	delta = complex(real(delta)/d, imag(delta)/d)
	i := start + delta
	log.Printf("Shortening: ]]]]]]]]]]]] : Searching start:%v, end:%v, delta:%v", start, end, delta)
	for ; i != end && pf.g.mat[i] == h; i += delta {
		connection = append(connection, i)
	}
	if i != end {
		log.Println("Shortening: ]]]]]]]]]]]] : Direct path with same height not found")
		return nil, false
	}
	connection = append(connection, end)
	return connection, len(connection) > 0
}
