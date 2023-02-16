package main

import (
	"day16/models"
	"day16/utility"
)

type Result struct {
	opened        []string
	totalReleased int
	timeElapsed   int
}

type PathFinder struct {
	rooms            map[string]*models.Room
	distances        map[Pair]int
	bitMask          *utility.BitMask[string]
	currentTimeLimit int
	listener         func(*Result)
}

func NewPathFinder(rooms map[string]*models.Room, distances map[Pair]int, bitMask *utility.BitMask[string]) *PathFinder {
	return &PathFinder{
		rooms:            rooms,
		distances:        distances,
		bitMask:          bitMask,
		currentTimeLimit: -1,
		listener:         nil,
	}
}

func (pf *PathFinder) SetTimeLimit(limit int) {
	pf.currentTimeLimit = limit
}

func (pf *PathFinder) SetListener(listener func(*Result)) {
	pf.listener = listener
}
