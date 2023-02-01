package models

import (
	"fmt"
	"math"
)

type Location complex128

func (l Location) TuningFrequency() int64 {
	return int64(real(l)*4000000 + imag(l))
}

func (l Location) Distance(other Location) float64 {
	delta := other - l
	return math.Abs(real(delta)) + math.Abs(imag(delta))
}

type Sensor struct {
	location              Location
	closestBeaconLocation Location
	closestBeaconDistance float64
}

func (s *Sensor) String() string {
	return fmt.Sprintf("Sensor%v", s.location)
}

func (s *Sensor) BeaconPossibleAt(l Location) bool {
	return s.location.Distance(l) > s.closestBeaconDistance
}

func NewSensor(l, b Location) *Sensor {
	return &Sensor{
		location:              l,
		closestBeaconLocation: b,
		closestBeaconDistance: l.Distance(b),
	}
}

type Beacon struct {
	l Location
}

func (b *Beacon) BeaconPossibleAt(l Location) bool {
	return b.l != l
}

func (b *Beacon) String() string {
	return fmt.Sprintf("Beacon%v", b.l)
}

func NewBeacon(l Location) *Beacon {
	return &Beacon{l}
}

type Object interface {
	BeaconPossibleAt(Location) bool
}
