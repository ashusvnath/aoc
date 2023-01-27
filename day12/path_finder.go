package main

import (
	. "day12/models"
	. "day12/utility"
	"log"
)

type PathFinder struct {
	q Queue[*Path]
	g *Grid
}

func (pf *PathFinder) FindPath(start complex128) *Path {
	g := pf.g
	pf.q.Enqueue(NewPath(g, start))
	visited := NewSet[complex128]()
	dropped := 0
	for pf.q.Len() > 0 {
		path := pf.q.Dequeue()
		currentIdx := path.CurrentLocation()
		if currentIdx == g.End() {
			log.Printf("Path found for %v: len: %d", start, path.Len())
			pf.q.Clear()
			return path
		}
		neighbours := g.Neighbours(currentIdx)
		for _, n := range neighbours {
			if !visited.Contains(n) && g.HeightAt(n)-g.HeightAt(path.CurrentLocation()) <= 1 {
				newPath := path.Append(n)
				visited.Add(n)
				pf.q.Enqueue(newPath)
			} else {
				dropped++
			}
		}
	}
	log.Printf("No path found for %v", start)
	pf.q.Clear()
	return nil
}

func (pf *PathFinder) FindHikingTrail(startTrail *Path) *Path {
	trail := *startTrail
	starts := pf.g.GetLocationsAtHeight(0).AsSlice()
	log.Printf("Searching %d starting points for shortest path to peak", len(starts))
	for _, idx := range starts {
		if idx == pf.g.Start() {
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
