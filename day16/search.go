package main

import (
	"day16/models"
	"day16/utility"
	"errors"
	"log"
	"sort"
)

func (pf *PathFinder) Search(start string, visited *utility.Set[string], timeRemaining int, releasedSoFar int, pathSoFar []string, who string) (*Result, string, error) {
	if len(pathSoFar) > 1 {
		pf.cachedResults[who] = append(pf.cachedResults[who], &Result{
			opened:        pathSoFar,
			timeElapsed:   pf.currentTimeLimit - timeRemaining,
			totalReleased: releasedSoFar,
		})
	}
	roomsToVisit := []string{}
	for _, v := range models.RoomIDsWithNonZeroReleaseRate {
		if !visited.Contains(v) {
			roomsToVisit = append(roomsToVisit, v)
		}
	}
	maxReleased := releasedSoFar
	var pathToMaxRelease = pathSoFar
	log.Printf("DFS: start: %s, visited: %v, timeremaining: %2d, released: %6d.", start, visited, timeRemaining, releasedSoFar)
	log.Printf("DFS: pathSoFar :%v, roomsRemaining:%v.", pathSoFar, roomsToVisit)
	if len(roomsToVisit) == 0 {
		returnValue := &Result{
			opened:        pathSoFar,
			totalReleased: releasedSoFar,
			timeElapsed:   pf.currentTimeLimit - timeRemaining,
		}
		//pf.cacheResult(returnValue, who)
		return returnValue, who, nil
	}
	if len(roomsToVisit) == 1 {
		target := roomsToVisit[0]
		pathFromStartToTarget := pf.distances[Pair{start, target}]
		timeTakenToVisitAndOpenLastValve := pathFromStartToTarget + 1
		newTimeRemaining := timeRemaining - timeTakenToVisitAndOpenLastValve
		if newTimeRemaining <= 0 {
			log.Printf(">>>>>>>> DFS: Ignoring %s at last level because time to travel %d.", target, timeTakenToVisitAndOpenLastValve)
			return nil, who, errors.New("cannot use this path")
		}
		releasedByTargetInRemainingTime := newTimeRemaining * models.AllRooms[target].ReleaseRate()
		returnPath := append([]string{}, pathSoFar...)
		returnPath = append(returnPath, target)
		returnValue := &Result{
			opened:        returnPath,
			totalReleased: releasedSoFar + releasedByTargetInRemainingTime,
			timeElapsed:   pf.currentTimeLimit - newTimeRemaining,
		}
		//pf.cacheResult(returnValue, who)
		return returnValue, who, nil
	}

	bestResult := &Result{totalReleased: releasedSoFar}

	for _, target := range roomsToVisit {
		pathFromStartToTarget := pf.distances[Pair{start, target}]
		timeTakenToVisitAndOpenValve := pathFromStartToTarget + 1
		newTimeRemaining := timeRemaining - timeTakenToVisitAndOpenValve
		if timeRemaining <= 0 {
			continue
		}
		releasedByTargetInRemainingTime := newTimeRemaining * models.AllRooms[target].ReleaseRate()
		visited.Add(target)
		forwardPath := append([]string{}, pathSoFar...)
		forwardPath = append(forwardPath, target)
		returnedResult, _, err := pf.Search(target, visited, newTimeRemaining, releasedSoFar+releasedByTargetInRemainingTime, forwardPath, who)
		if err != nil {
			visited.Remove(target)
			continue
		}
		if returnedResult.totalReleased > bestResult.totalReleased {
			bestResult = returnedResult
		}
		visited.Remove(target)
	}
	log.Printf("DFS: **************** path: %v, released: %5d", pathToMaxRelease, maxReleased)
	//pf.cacheResult(bestResult, who)
	return bestResult, who, nil
}

func (pf *PathFinder) SearchWithCache(start string, maxTime int, who string) ([]*Result, error) {
	selectedKeys := []string{}
	for k, v := range pf.cachedResults {
		if k != who {
			selectedKeys = append(selectedKeys, k)
		}
		sort.Slice(v, func(left, right int) bool {
			return v[left].timeElapsed  < v[right].timeElapsed
		})
	}
	bestTotal := -1
	bestResults := []*Result{{}, {}}
	for _, k := range selectedKeys {
		for _, cachedResult := range pf.cachedResults[k] {
			if cachedResult.timeElapsed > maxTime {
				continue
			}
			log.Printf("SearchWithCache: Examining %v", cachedResult)
			visited := utility.NewSet[string]()
			visited.AddAll(cachedResult.opened...)
			newResult, _, err := pf.Search(start, visited, maxTime, 0, []string{start}, who)
			if err != nil {
				log.Fatalf("SearchWithCache: error when searching %v", err)
			}
			if newResult.totalReleased+cachedResult.totalReleased > bestTotal {
				bestTotal = newResult.totalReleased + cachedResult.totalReleased
				bestResults = []*Result{cachedResult, newResult}
			}
		}
	}
	return bestResults, nil
}

func (pf *PathFinder) FindPath(from, to string) []string {
	q := utility.NewQueue[[]string]()
	visited := utility.NewSet[string]()
	q.Enqueue([]string{from})
	for !q.IsEmpty() {
		current := q.Dequeue()
		l := len(current)
		for _, n := range pf.rooms[current[l-1]].ConnectedRoomIds() {
			if visited.Contains(n) {
				continue
			}
			next := append([]string{}, current...)
			next = append(next, n)
			if n == to {
				return next
			}
			q.Enqueue(next)
			visited.Add(n)
		}
	}
	return nil
}
