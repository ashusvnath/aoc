package models

import (
	"day15/utility"
	"log"
	"math"
)

type SensorProximityObserverForY struct {
	target                   float64
	forbiddenBeaconLocations *utility.Set[Location]
	knownBeaconLocations     *utility.Set[Location]
}

type SensorProximityObserver interface {
	Listen(sensor *Sensor)
	Count() int
}

func (s SensorProximityObserverForY) Listen(sensor *Sensor) {
	sensorLocation := sensor.location
	distanceToSensor := math.Abs(imag(sensorLocation) - s.target)
	sensorX := real(sensorLocation)
	xRange := sensor.closestBeaconDistance - distanceToSensor
	log.Printf("Sensor at %4v is %2.0f units away from closest sensor and %2.0f units from line y=%2.0f",
		sensorLocation, sensor.closestBeaconDistance, distanceToSensor, s.target)
	if imag(sensor.closestBeaconLocation) == s.target {
		s.forbiddenBeaconLocations.Remove(sensor.closestBeaconLocation)
		log.Printf("Known beacon locations updated to %d", s.knownBeaconLocations.Len())
		s.knownBeaconLocations.Add(sensor.closestBeaconLocation)
	}
	if xRange <= 0 {
		return
	}
	for x := sensorX - xRange; x <= sensorX+xRange; x++ {
		improbableBeaconLocation := Location(complex(x, s.target))
		if s.knownBeaconLocations.Contains(improbableBeaconLocation) {
			continue
		}
		s.forbiddenBeaconLocations.Add(improbableBeaconLocation)
	}
}

func (s SensorProximityObserverForY) Count() int {
	return s.forbiddenBeaconLocations.Len()
}

func NewSensorProximityObserver(y float64) SensorProximityObserver {
	return &SensorProximityObserverForY{
		target:                   y,
		forbiddenBeaconLocations: utility.NewSet[Location](),
		knownBeaconLocations:     utility.NewSet[Location](),
	}
}
