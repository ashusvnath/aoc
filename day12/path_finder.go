package main

import (
	"log"
)

type PathFinder struct {
	q *Queue[*Path]
	g *Grid
}

func (pf *PathFinder) FindPath(start complex128) *Path {
	g := pf.g
	pf.q.Push(NewPath(g, start))
	visited := NewSet[complex128]()
	dropped := 0
	for !pf.q.IsEmpty() {
		path := pf.q.Pop()
		currentIdx := path.currentIdx
		if currentIdx == g.end {
			log.Printf("Path found for %v: len: %d", start, path.Len())
			pf.q.Clear()
			return path
		}
		neighbours := g.Neighbours(currentIdx)
		for _, n := range neighbours {
			if !visited.Contains(n) && g.mat[n]-g.mat[path.currentIdx] <= 1 {
				newPath := path.Add(n)
				visited.Add(n)
				pf.q.Push(newPath)
			} else {
				dropped++
			}
		}
	}
	log.Printf("No path found for %v", start)
	pf.q.Clear()
	return nil
}

func (pf *PathFinder) HikingTrail(startTrail *Path) *Path {
	trail := *startTrail
	starts := pf.g.idxsByHeight[0].AsSlice()
	log.Printf("Searching %d starting points for shortest", len(starts))
	for _, idx := range starts {
		if idx == pf.g.start {
			continue
		}
		path := pf.FindPath(idx)
		if path != nil && path.Len() < trail.Len() {
			trail = *path
		}
	}
	return &trail
}

func NewPathFinder(g *Grid) *PathFinder {
	return &PathFinder{
		q: NewQueue[*Path](),
		g: g,
	}
}
