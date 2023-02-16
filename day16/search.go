package main

import (
	"day16/models"
	"day16/utility"
	"errors"
	"log"
)

func (pf *PathFinder) Search(start string, visited utility.Set[string], timeRemaining int, releasedSoFar int, pathSoFar []string, who string) (*Result, string, error) {
	if len(pathSoFar) > 1 && pf.listener != nil {
		pf.listener(&Result{
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
		return returnValue, who, nil
	}

	bestResult := &Result{totalReleased: releasedSoFar, opened: pathSoFar, timeElapsed: pf.currentTimeLimit - timeRemaining}

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
	return bestResult, who, nil
}
