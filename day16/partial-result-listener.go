package main

import (
	"day16/models"
	"day16/utility"
	"fmt"
	"log"
)

type ParialResultListener struct {
	distances  map[Pair]int
	maxTime    int
	bestResult [2]*Result
	pf         *PathFinder
	bestTotal  int
	count      int64
}

func NewParialResultListener(distances map[Pair]int, maxTime int, bitMask *utility.BitMask[string]) *ParialResultListener {
	pf := NewPathFinder(models.AllRooms, distances, bitMask)
	pf.SetTimeLimit(maxTime)
	return &ParialResultListener{
		distances:  distances,
		maxTime:    maxTime,
		pf:         pf,
		bestResult: [2]*Result{nil, nil},
	}
}

func (prl *ParialResultListener) Listen(result *Result) {
	if result.timeElapsed > prl.maxTime || len(result.opened) < len(models.RoomIDsWithNonZeroReleaseRate)/2-1 {
		return
	}
	prl.count++
	if prl.count%1000 == 0 {
		fmt.Print(".")
	}

	visited := utility.NewBitMaskSet(prl.pf.bitMask)
	visited.AddAll(result.opened...)
	otherResult, _, err := prl.pf.Search("AA", visited, prl.maxTime, 0, []string{"AA"}, "elephant")
	if err != nil {
		log.Printf("PartialResultListener: error when listening to result %v: %v", *result, err)
	}
	if prl.bestResult[0] == nil {
		prl.bestResult = [2]*Result{prl.RecalculateTotal(result), otherResult}
		prl.bestTotal = prl.bestResult[0].totalReleased + prl.bestResult[1].totalReleased
	} else {
		newResult := prl.RecalculateTotal(result)
		latestResultTotal := newResult.totalReleased + otherResult.totalReleased
		if prl.bestTotal < latestResultTotal {
			prl.bestResult = [2]*Result{newResult, otherResult}
			prl.bestTotal = latestResultTotal
		}
	}
	prl.count++
	if prl.count%1000 == 0 {
		log.Printf("PRL: iteration: %5d, bestTotalSoFar: %5d,  Results:%v, %v\n", prl.count, prl.bestTotal, prl.bestResult[0], prl.bestResult[1])
	}
}

func (prl *ParialResultListener) BestResult() ([]*Result, int) {
	return prl.bestResult[:], prl.bestTotal
}

func (prl *ParialResultListener) RecalculateTotal(result *Result) *Result {
	newTotal := 0
	timeElapsed := 0
	for idx, to := range result.opened[1:] {
		from := result.opened[idx]
		d := prl.pf.distances[Pair{from, to}]
		timeElapsed += (d + 1)
		newTotal += models.AllRooms[to].ReleaseRate() * (prl.maxTime - timeElapsed)
	}
	return &Result{result.opened, newTotal, result.timeElapsed}
}
