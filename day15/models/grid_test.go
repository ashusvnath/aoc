package models

import (
	"day15/assert"
	"log"
	"testing"
)

func TestGrid(t *testing.T) {
	t.Run("Should notify registered listeners when new sensor added to grid", func(t *testing.T) {
		sLoc := Location(1 + 10i)
		bLoc := Location(10 + 11i)
		s1 := NewSensor(sLoc, bLoc)
		notificationCount := 0
		n := func(in *Sensor) {
			notificationCount++
			log.Printf("Recieved notification for sensor %v", in)
			assert.Equal(s1.location, in.location, t)
			assert.Equal(s1.closestBeaconLocation, in.closestBeaconLocation, t)
			assert.Equal(s1.closestBeaconDistance, in.closestBeaconDistance, t)
		}

		g := NewGrid()
		g.Register(n)

		g.AddSensor(sLoc, bLoc)
		s1 = NewSensor(sLoc+1+1i, bLoc)
		g.AddSensor(sLoc+1+1i, bLoc)

		s1 = NewSensor(sLoc-1-1i, bLoc)
		g.AddSensor(sLoc-1-1i, bLoc)

		assert.Equal(3, notificationCount, t)
	})
}
