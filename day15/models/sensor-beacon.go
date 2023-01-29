package models

import (
	"math"
)

type Location complex128

func (l Location) Distance(other Location) float64 {
	delta := other - l
	return math.Abs(real(delta)) + math.Abs(imag(delta))
}

type Type int

const (
	SENSOR Type = 1
	BEACON Type = 2
)

type Sensor struct {
	location              Location
	closestBeaconLocation Location
	closestBeaconDistance float64
}

func (s *Sensor) Type() Type {
	return SENSOR
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

func (b *Beacon) Type() Type {
	return BEACON
}

func NewBeacon(l Location) *Beacon {
	return &Beacon{l}
}

type Object interface {
	Type() Type
}
