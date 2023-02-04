package main

import (
	"day16/models"
	"day16/utility"
	"errors"
	"log"
)

func Search(start string, visited *utility.Set[string], timeRemaining int, releasedSoFar int, pf PathFunc, pathSoFar []string) (int, []string, error) {
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
	if len(roomsToVisit) < 2 {
		target := roomsToVisit[0]
		pathFromStartToTarget := pf(start, target)
		timeTakenToVisitAndOpenLastValve := pathFromStartToTarget + 1
		newTimeRemaining := timeRemaining - timeTakenToVisitAndOpenLastValve
		if newTimeRemaining <= 0 {
			log.Printf(">>>>>>>> DFS: Ignoring %s at last level because time to travel %d.", target, timeTakenToVisitAndOpenLastValve)
			return -1, nil, errors.New("cannot use this path")
		}
		releasedByTargetInRemainingTime := newTimeRemaining * models.AllRooms[target].ReleaseRate()
		returnPath := append([]string{}, pathSoFar...)
		returnPath = append(returnPath, target)
		return releasedSoFar + releasedByTargetInRemainingTime, returnPath, nil
	}

	for _, target := range roomsToVisit {
		pathFromStartToTarget := pf(start, target)
		timeTakenToVisitAndOpenValve := pathFromStartToTarget + 1
		newTimeRemaining := timeRemaining - timeTakenToVisitAndOpenValve
		if timeRemaining <= 0 { 
			continue
		}
		releasedByTargetInRemainingTime := newTimeRemaining * models.AllRooms[target].ReleaseRate()
		visited.Add(target)
		forwardPath := append([]string{}, pathSoFar...)
		forwardPath = append(forwardPath, target)
		returnedTotalRelease, returnedPath, err := Search(target, visited, newTimeRemaining, releasedSoFar+releasedByTargetInRemainingTime, pf, forwardPath)
		if err != nil {
			visited.Remove(target)
			continue
		}
		if returnedTotalRelease > maxReleased {
			maxReleased = returnedTotalRelease
			pathToMaxRelease = returnedPath
		}
		visited.Remove(target)
	}
	log.Printf("DFS: **************** path: %v, released: %5d", pathToMaxRelease, maxReleased)
	return maxReleased, pathToMaxRelease, nil
}

