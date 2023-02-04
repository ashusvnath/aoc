package main

import (
	"day16/utility"
	"fmt"
)

type PathFunc func(start, end string) int

// func Part1(bfs *utility.BFS[string]) {
func Part1(pds map[Pair]int) {
	pathFunc := func(from, to string) int {
		return pds[Pair{from, to}] //exlude start in path returned
	}
	visited := utility.NewSet[string]()
	totalReleased, path, _ := Search("AA", visited, 30, 0, pathFunc, []string{"AA"})
	fmt.Printf("Part1: %d, %v", totalReleased, path)
	// q := utility.NewQueue[BestPathNode]()
	// s := utility.NewStack[utility.Queue[BestPathNode]]()
	// initQ := utility.NewQueue[BestPathNode]()
	// initQ.Enqueue(BestPathNode{"AA", 0, 30})
	// visited := utility.NewSet[string]()
	// s.Push(initQ)
	// for s.Len() > 0 {
	// 	bpnSource := s.Peek()
	// 	bpn := bpnSource.Peek()
	// 	timeRemaining := bpn.remainingTime
	// 	start := bpn.id
	// 	visited.Add(start)
	// 	next := FindNextBest(start, timeRemaining, visited, pathFunc)
	// 	for {
	// 		s.Push(q)
	// 		q = FindNextBest(next.id, next.remainingTime, visited, pathFunc)
	// 	}
	// 	s.Peek().Dequeue()
	// }
}

// type BestPathNode struct {
// 	id            string
// 	released      int
// 	remainingTime int
// }

// func FindNextBest(start string, timeLeft int, visited *utility.Set[string], pathFunc func(start, end string) []string) utility.Queue[BestPathNode] {
// 	result := utility.NewQueue[BestPathNode]()
// 	nodes := []BestPathNode{}
// 	for _, roomId := range models.RoomIDsWithNonZeroReleaseRate {
// 		if visited.Contains(roomId) {
// 			continue
// 		}
// 		path := pathFunc(start, roomId)
// 		timeAfterPathTakenAndValveOpened := timeLeft - len(path) - 1
// 		if timeAfterPathTakenAndValveOpened <= 0 {
// 			continue
// 		}
// 		eventualRelease := timeAfterPathTakenAndValveOpened * models.AllRooms[roomId].ReleaseRate()
// 		nodes := append(nodes, BestPathNode{roomId, eventualRelease, timeAfterPathTakenAndValveOpened})
// 		//fmt.Printf("%6s, %3d, %4d, %v\n", roomId, len(path)+1, eventualRelease, path[1:])
// 		sort.Slice(nodes, func(left, right int) bool {
// 			return nodes[left].remainingTime > nodes[right].remainingTime
// 		})
// 		for _, node := range nodes {
// 			result.Enqueue(node)
// 		}
// 	}
// 	return result
// }

// func DFS(bfs *utility.BFS[string], elapsedTime, projectedTotalFlow int, start string, visited *utility.Set[string], soFar []string) (int, []string) {
// 	oldSoFar := make([]string, len(soFar))
// 	copy(oldSoFar, soFar)
// 	var bestPath []string
// 	maxEstimatedFlow := projectedTotalFlow + GetTotalFlow(visited.AsSlice())*(30-elapsedTime)
// 	for _, roomId := range models.RoomIDsWithNonZeroReleaseRate {
// 		if visited.Contains(roomId) {
// 			continue
// 		}
// 		p := bfs.FindShortestPath(start, roomId, func(in string) []string {
// 			return models.AllRooms[in].ConnectedRoomIds()
// 		})
// 		timeToSelectedRoomAndOpen := len(p)
// 		if elapsedTime+timeToSelectedRoomAndOpen >= 30 {
// 			continue
// 		}
// 		newSoFar := append(oldSoFar, p...)
// 		newEstimatedTotal := projectedTotalFlow + GetTotalFlow(visited.AsSlice())*(30-elapsedTime-timeToSelectedRoomAndOpen)
// 		visited.Add(roomId)
// 		resultantFlow, path := DFS(bfs, elapsedTime+timeToSelectedRoomAndOpen, newEstimatedTotal, roomId, visited, newSoFar)
// 		if resultantFlow > maxEstimatedFlow {
// 			maxEstimatedFlow = resultantFlow
// 			bestPath = append([]string{}, soFar...)
// 			bestPath = append(bestPath, p...)
// 			bestPath = append(bestPath, path...)
// 		}
// 		visited.Remove(roomId)
// 	}
// 	return maxEstimatedFlow, bestPath
// }

// func GetTotalFlow(selectedRooms []string) int {
// 	total := 0
// 	for _, roomId := range selectedRooms {
// 		total += models.AllRooms[roomId].ReleaseRate()
// 	}
// 	return total
// }
