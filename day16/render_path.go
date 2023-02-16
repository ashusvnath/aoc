package main

import (
	"day16/models"
	"fmt"
	"strings"
)

func RenderPath(path []string, pf *PathFinder, timeRemaining int) (string, int, int) {
	sb := strings.Builder{}
	timeTaken := 0
	totalReleased := 0
	totalReleaseRate := 0
	for idx, to := range path[1:] {
		from := path[idx]
		subPath := pf.FindPath(from, to)
		timeTaken += len(subPath)
		totalReleaseRate += models.AllRooms[to].ReleaseRate()
		totalReleased += models.AllRooms[to].ReleaseRate() * (timeRemaining - timeTaken)
		sb.WriteString(strings.Join(subPath, " -> "))
		sb.WriteString(fmt.Sprintf("(time:%2d(%2d), released:%6d, rate: %d)\n", timeTaken, len(subPath), totalReleased, totalReleaseRate))
	}
	return sb.String(), totalReleased, timeTaken
}
