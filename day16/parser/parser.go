package parser

import (
	"day16/models"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	roomDetailPatternString = `Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? ([A-Z, ]{2,})`
)

func Parse(input string) *models.Room {
	roomDetailPattern := regexp.MustCompile(roomDetailPatternString)
	matches := roomDetailPattern.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		log.Printf("Match %v", match[1:])
		id := match[1]
		rate, _ := strconv.Atoi(match[2])
		connectedIds := strings.Split(match[3], ", ")
		models.NewRoom(id, rate, connectedIds...)
	}

	return models.AllRooms[matches[0][1]]
}
