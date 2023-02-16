package main

import (
	"day16/models"
)

type Result struct {
	opened        []string
	totalReleased int
	timeElapsed   int
}

type PathFinder struct {
	rooms            map[string]*models.Room
	distances        map[Pair]int
	cachedResults    map[string][]*Result
	currentTimeLimit int
}

func NewPathFinder(rooms map[string]*models.Room, distances map[Pair]int) *PathFinder {
	return &PathFinder{
		rooms:            rooms,
		distances:        distances,
		cachedResults:    make(map[string][]*Result),
		currentTimeLimit: -1,
	}
}

func (pf *PathFinder) SetTimeLimit(limit int) {
	pf.currentTimeLimit = limit
}
