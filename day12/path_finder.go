package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"sort"
)

type PathFinder struct {
	visited          Set[int]
	path             []int
	pi               *ParsedInput
	knownDistances   map[int]int
	backtrackedNodes Set[int]
}

func (pf *PathFinder) FindPath() {
	pi := pf.pi
	currentIdx := pi.start
	currentHeight := 0
	pf.visited.Add(currentIdx)
	//pf.path = append(pf.path, currentIdx)
	for currentIdx != pi.end {
		prevIdx := currentIdx
		neighbours := pf.rankedNeighbours(currentIdx)
		log.Printf("idx: %v, height: %v, neighbours, :%#v", currentIdx, currentHeight, neighbours)
		for _, n := range neighbours {
			if !pf.visited.Contains(n[0]) {
				currentHeight = n[1]
				pf.visited.Add(n[0])
				currentIdx = n[0]
				pf.path = append(pf.path, n[0])
				break
			}
		}
		if currentIdx == prevIdx {
			//Backtracking because no progress can be made
			backtrackCount := 1
			if pf.backtrackedNodes.Contains(currentIdx) {
				log.Printf("Could be in a loop. Increasing backtrack count")
				backtrackCount++
			}
			l := len(pf.path)
			log.Printf("***************path so far: (%d). Backtracking", l) //, pf.path)
			//delete(pf.visited, currentIdx)
			pf.backtrackedNodes.Add(currentIdx)
			currentIdx = pf.path[l-1-backtrackCount]
			newPath := make([]int, len(pf.path)-1-backtrackCount)
			copy(newPath, pf.path)
			newPath = append(newPath, currentIdx)
			pf.path = newPath
		}
	}
	pf.Shorten()
}

func (pf *PathFinder) Shorten() {
	log.Printf("]]]]]]]]]]]] Starting shortening procedure")
	pi := pf.pi
	cols := pi.cols
	shortened := true
	shorteningRange := 4
	for shorteningRange < 5 {
		shortened = false
		prev := pi.start
		l := len(pf.path)
		dh := make([]int, l)
		dx := make([]int, l)
		dy := make([]int, l)
		for i := 0; i < len(pf.path); i++ {
			idx := pf.path[i]
			dh[i] = pi.GetByIdx(idx) - pi.GetByIdx(prev)
			dx[i] = (idx % cols) - (prev % cols)
			dy[i] = (idx / cols) - (prev / cols)

			if i > shorteningRange-1 {
				shortened = pf.tryShorten(shorteningRange, i, dh, dx, dy)
				if shortened {
					break
				}
			}
			prev = idx
		}
		if !shortened {
			shorteningRange++
		}
	}
}

func (pf *PathFinder) tryShorten(shorteningRange int, i int, dh, dx, dy []int) bool {
	shortened := false
	dh_range := 0
	dx_range := 0
	dy_range := 0
	dh_range_sq := 0
	dx_range_sq := 0
	dy_range_sq := 0
	for j := 0; j < shorteningRange-1; j++ {
		dh_range += dh[i-j]
		dh_range_sq += dh[i-j] * dh[i-j]
		dx_range += dx[i-j]
		dx_range_sq += dx_range * dx_range
		dy_range += dy[i-j]
		dy_range_sq += dy_range * dy_range
	}
	dh_range += dh[i-shorteningRange]
	dh_range_sq += dh[i-shorteningRange] * dh[i-shorteningRange]
	if ((dy_range_sq != 0 && dy_range == 0 && dx_range%2 == 1) || (dx_range == 0 && dx_range_sq != 0 && dy_range%2 == 1)) && (dh_range_sq == 0) {
		newPath := []int{}
		connection, success := pf.createPath(pf.path[i-shorteningRange+1], pf.path[i])
		if !success {
			return false
		}
		log.Printf("Shortening required at: %4d: %#v", pf.path[i-shorteningRange+1], pf.IdxAsRC(pf.path[i-shorteningRange:i+3]))
		log.Printf("Shortcut for range %d found between: %4d, %4d: %2d: %#v", shorteningRange, pf.path[i], pf.path[i-shorteningRange+1], len(connection), pf.IdxAsRC(connection))
		newPath = append(newPath, pf.path[:i-shorteningRange]...)
		newPath = append(newPath, connection...)
		newPath = append(newPath, pf.path[i+1:]...)
		pf.path = newPath
		shortened = true
	}
	return shortened
}

func (pf *PathFinder) createPath(idx1, idx2 int) ([]int, bool) {
	connection := []int{}
	cols := pf.pi.cols
	h := pf.pi.GetByIdx(idx1)
	r1, r2, c1, c2 := idx1/cols, idx2/cols, idx1%cols, idx2%cols
	if r1 == r2 {
		for i := c1; i <= c2 && pf.pi.GetByRC(r1, i) == h; i++ {
			connection = append(connection, r1*cols+i)
		}
	} else if c1 == c2 {
		for i := r1; i <= r2 && pf.pi.GetByRC(i, c1) == h; i++ {
			connection = append(connection, i*cols+c1)
		}
	} else {
		log.Printf("XXXXXXXXXXXX : Cannot shorten no row or column match")
	}
	return connection, len(connection) > 0
}

func (pf *PathFinder) IdxAsRC(in []int) []string {
	cols := pf.pi.cols
	result := []string{}
	for _, idx := range in {
		h := pf.pi.GetByIdx(idx)
		result = append(result, fmt.Sprintf("(%3d,%3d):%2d", idx/cols, idx%cols, h))
	}
	return result

}

func (pf *PathFinder) VisualizePath() string {
	pi := pf.pi
	path := make([]int, len(pf.path))
	copy(path, pf.path)
	path = append(path, pi.end)
	displayData := []rune(string(pf.pi.rawData))
	displayByteMap := map[int]rune{
		-1: '\u2190', 1: '\u2192',
		(0 - pi.cols): '\u2191', (pi.cols): '\u2193',
		0: 'E',
	}
	for x, idx := range pf.path {
		row, col := idx/pi.cols, idx%pi.cols
		displayIdx := row*(pi.cols+1) + col
		nextEntry := path[x+1]
		d := displayByteMap[nextEntry-idx]
		if d == 0 {
			d = '?'
		}
		displayData[displayIdx] = d
	}
	return regexp.MustCompile(`[bd-z]`).ReplaceAllString(string(displayData), "·")
}

func (pf *PathFinder) dist(idx1, idx2, cols int) int {
	key := idx1 * idx2
	if d := pf.knownDistances[key]; d != 0 {
		return d
	}
	r1, c1 := idx1/cols, idx1%cols
	r2, c2 := idx2/cols, idx2%cols
	dx := math.Abs(float64(r2 - r1))
	dy := math.Abs(float64(c2 - c1))
	v := int(dx + dy)
	pf.knownDistances[key] = v
	return v
}

func (pf *PathFinder) rankedNeighbours(idx int) [][2]int {
	pi := pf.pi
	currentHeight := pi.GetByIdx(idx)
	_n := pi.Neighbours(idx)
	log.Printf("Unranked neighbours of %d: %#v", idx, _n)
	n := pf.filter(_n, currentHeight)
	sort.Slice(n, func(i, j int) bool {
		//Rank first by height
		if n[i][1] > n[j][1] {
			return true
		} else if n[i][1] < n[j][1] {
			return false
		}
		//Rank by distance to end if height same
		return pf.dist(n[i][0], pi.end, pi.cols) < pf.dist(n[j][0], pi.end, pi.cols)
	})
	return n
}

func (pf *PathFinder) filter(in []int, h int) [][2]int {
	result := [][2]int{}
	for _, idx := range in {
		nHeight := pf.pi.GetByIdx(idx)
		if nHeight-h > 1 || pf.visited.Contains(idx) {
			continue
		}
		result = append(result, [2]int{idx, nHeight})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][1] > result[j][1]
	})
	log.Printf("Filtered neigbhours %#v", result)
	return result
}

func NewPathFinder(pi *ParsedInput) *PathFinder {
	return &PathFinder{
		visited:          make(Set[int]),
		path:             []int{},
		pi:               pi,
		knownDistances:   make(map[int]int),
		backtrackedNodes: make(Set[int]),
	}
}
