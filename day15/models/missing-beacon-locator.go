package models

import (
	"log"
	"math"
)

type MissingBeaconLocator struct {
	scanLines map[float64]*ScanLine
	size      float64
}

func NewMissingBeaconLocator(size float64) *MissingBeaconLocator {
	return &MissingBeaconLocator{
		scanLines: make(map[float64]*ScanLine),
		size:      size,
	}
}

func (mbl *MissingBeaconLocator) Process(sensor *Sensor) {
	mbl.processBeaconLocation(sensor.closestBeaconLocation)
	mbl.processSensor(sensor.location, sensor.closestBeaconDistance)
}

func (mbl *MissingBeaconLocator) processBeaconLocation(beaconLocation Location) {
	beaconY := imag(beaconLocation)
	if beaconY < 0 || beaconY > mbl.size+1 {
		return
	}
	beaconX := real(beaconLocation)
	if mbl.scanLines[beaconY] == nil {
		mbl.scanLines[beaconY] = NewScanLine(float64(mbl.size), beaconY)
	}
	mbl.scanLines[beaconY].Exclude(beaconX, beaconX)
}

func (mbl *MissingBeaconLocator) processSensor(sensorLocation Location, closestBeaconDistance float64) {
	sensorY := imag(sensorLocation)
	sensorX := real(sensorLocation)

	idxStart := math.Max(float64(sensorY)-closestBeaconDistance, 0)
	idxStop := math.Min(float64(sensorY)+closestBeaconDistance, mbl.size+1)
	log.Printf("MBL.Process: Updating scanlines y=%.0f to %.0f", idxStart, idxStop)
	for idx := idxStart; idx <= idxStop; idx += 1.0 {
		scanLine := mbl.scanLines[idx]
		delta := closestBeaconDistance - math.Abs(sensorY-idx)
		if scanLine == nil {
			scanLine = NewScanLine(mbl.size, idx)
			mbl.scanLines[idx] = scanLine
		}
		windowStart := sensorX - delta
		windowEnd := sensorX + delta
		scanLine.Exclude(windowStart, windowEnd)
	}
}

func (mbl *MissingBeaconLocator) Locate(g *Grid) Location {
	for y, scanLine := range mbl.scanLines {
		for _, xWindow := range scanLine.PossibleBeaconLocations() {
			x := xWindow[0] - 1
			if x < 0 {
				continue
			}
			loc := Location(complex(x, y))
			if g.BeaconPossibleAt(loc) {
				return loc
			}

			x = xWindow[0] + 1
			if x > mbl.size {
				continue
			}
			loc = Location(complex(x, y))
			if g.BeaconPossibleAt(loc) {
				return loc
			}

			x = xWindow[1] - 1
			if x < 0 {
				continue
			}
			loc = Location(complex(x, y))
			if g.BeaconPossibleAt(loc) {
				return loc
			}

			x = xWindow[1] + 1
			if x > mbl.size {
				continue
			}
			loc = Location(complex(x, y))
			if g.BeaconPossibleAt(loc) {
				return loc
			}
		}
	}
	return Location(complex(-1, -1))
}
