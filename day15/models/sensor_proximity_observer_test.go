package models

import (
	"day15/assert"
	"testing"
)

func TestSensorProximityObserver(t *testing.T) {
	t.Run("constructor should return new proximity observer with specified target y co-ordinate", func(t *testing.T) {
		target := float64(20)
		spofy := GetSPOfYWithTarget(target, t)
		assert.Equal(target, spofy.target, t)
	})

	t.Run("SPOfY.Listen should process incoming sensor data when empty", func(t *testing.T) {
		target := float64(20)
		spofy := GetSPOfYWithTarget(target, t)
		sLoc := Location(complex(10, 10))
		bLoc := Location(complex(10, 15))
		sensor := NewSensor(sLoc, bLoc)

		spofy.Listen(sensor)

		assert.Equal(0, spofy.Count(), t)
		assert.Equal(0, spofy.knownBeaconLocations.Len(), t)
		assert.Equal(0, spofy.forbiddenBeaconLocations.Len(), t)
	})

	t.Run("SPOfY.Listen should process incoming sensor data", func(t *testing.T) {
		target := float64(13)
		spofy := GetSPOfYWithTarget(target, t)
		sLoc := Location(complex(10, 10))
		bLoc := Location(complex(10, 15))
		sensor := NewSensor(sLoc, bLoc)

		spofy.Listen(sensor)

		assert.Equal(0, spofy.knownBeaconLocations.Len(), t)
		assert.Equal(5, spofy.forbiddenBeaconLocations.Len(), t)
	})

	t.Run("SPOfY.Listen should process incoming sensor data when the observer has a clashing forbidden beacon location", func(t *testing.T) {
		target := float64(15)
		spofy := GetSPOfYWithTarget(target, t)
		sLoc := Location(complex(10, 10))
		bLoc := Location(complex(10, 15))

		//Setup clashing forbidden beacon location on observer
		spofy.forbiddenBeaconLocations.Add(bLoc)
		spofy.forbiddenBeaconLocations.Add(bLoc + 1)

		//Before sensor data added
		assert.Equal(0, spofy.knownBeaconLocations.Len(), t)
		assert.Equal(2, spofy.Count(), t)

		//Add sensor data
		spofy.Listen(NewSensor(sLoc, bLoc))

		//After sensor data added
		assert.Equal(1, spofy.knownBeaconLocations.Len(), t)
		assert.Equal(1, spofy.Count(), t)
	})

	t.Run("SPOfY.Listen should skip adding known beacon locations to forbidden beacon location", func(t *testing.T) {
		target := float64(12)
		spofy := GetSPOfYWithTarget(target, t)
		sLoc := Location(complex(10, 10))
		bLoc := Location(complex(10, 15))

		//Setup known beacon location on observer
		spofy.knownBeaconLocations.Add(Location(complex(11, 12)))

		//Before sensor data added
		assert.Equal(0, spofy.Count(), t)

		//Add sensor data
		spofy.Listen(NewSensor(sLoc, bLoc))

		//After sensor data added
		assert.Equal(1, spofy.knownBeaconLocations.Len(), t)
		assert.Equal(6, spofy.Count(), t)
	})
}

func GetSPOfYWithTarget(target float64, t *testing.T) *SensorProximityObserverForY {
	spo := NewSensorProximityObserver(target)
	spofy, ok := spo.(*SensorProximityObserverForY)
	assert.True(ok, t)
	return spofy
}
