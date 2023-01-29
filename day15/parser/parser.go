package parser

import (
	"day15/models"
	"day15/utility"
	"regexp"
	"strconv"
	"strings"
)

var lineRegex = regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

func ParseLine(input string) (models.Location, models.Location) {
	matched := lineRegex.FindStringSubmatch(input)
	sensorX, _ := strconv.ParseFloat(matched[1], 64)
	sensorY, _ := strconv.ParseFloat(matched[2], 64)
	beaconX, _ := strconv.ParseFloat(matched[3], 64)
	beaconY, _ := strconv.ParseFloat(matched[4], 64)

	return models.Location(complex(sensorX, sensorY)), models.Location(complex(beaconX, beaconY))
}

func Parse(input string, obs utility.Notifiable[*models.Sensor]) *models.Grid {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	grid := models.NewGrid()
	grid.Register(obs)
	for _, line := range lines {
		sensor, beacon := ParseLine(line)
		grid.AddSensor(sensor, beacon)
	}
	return grid
}
