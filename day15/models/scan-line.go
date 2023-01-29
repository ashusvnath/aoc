package models

import (
	"math"
	"sort"
)

type window [2]float64

func (w window) lesserThan(other window) bool {
	return w[0] < other[0]
}

func (w window) overlapsOrTouches(other window) bool {
	if other.lesserThan(w) {
		return other.overlapsOrTouches(w)
	}
	return (w[1] >= other[0]-1)
}

func (w window) contains(other window) bool {
	return (w[0] <= other[0]) && (w[1] >= other[1])
}

func (w window) merge(other window) window {
	if other.lesserThan(w) {
		return other.merge(w)
	}
	return window{w[0], other[1]}
}

type ScanLine struct {
	exclusions    []window
	max           float64
	fullyExcluded bool
}

func (s *ScanLine) Size() float64 {
	return s.max
}

func (s *ScanLine) Exclude(start, stop float64) {
	if s.fullyExcluded {
		return
	}
	newWindowStart := math.Max(start, 0)
	newWindowStop := math.Min(s.max, stop)
	if newWindowStart > newWindowStop {
		return
	}
	newWindow := window{newWindowStart, newWindowStop}
	newExclusions := []window{}
	for i := 0; i < len(s.exclusions); i++ {
		oldWindow := s.exclusions[i]
		if newWindow.contains(oldWindow) {
			continue
		} else if newWindow.overlapsOrTouches(oldWindow) {
			newWindow = oldWindow.merge(newWindow)
		} else {
			newExclusions = append(newExclusions, oldWindow)
		}
	}
	newExclusions = append(newExclusions, newWindow)
	if len(newExclusions) > 1 {
		sort.SliceStable(newExclusions, func(i, j int) bool {
			return newExclusions[i].lesserThan(newExclusions[j])
		})
	} else if newWindow[0] == 0 && newWindow[1] == s.max {
		s.fullyExcluded = true
	}
	s.exclusions = newExclusions
}

func NewScanLine(f float64) *ScanLine {
	return &ScanLine{[]window{}, f, false}
}
