package models

import "day15/utility"

type Grid struct {
	mat map[Location]Object
	obs *utility.Observable[*Sensor]
}

func (g *Grid) ObjectAt(l Location) Object {
	return g.mat[l]
}

func (g *Grid) AddSensor(sensorLocation, beaconLocation Location) {
	distance := sensorLocation.Distance(beaconLocation)
	sensor := &Sensor{sensorLocation, beaconLocation, distance}
	g.mat[sensorLocation] = sensor
	g.mat[beaconLocation] = &Beacon{beaconLocation}
	g.obs.NotifyWith(sensor)
}

func (g *Grid) Register(n utility.Notifiable[*Sensor]) {
	g.obs.Register(n)
}

func NewGrid() *Grid {
	return &Grid{
		mat: make(map[Location]Object),
		obs: utility.NewObservable[*Sensor](nil),
	}
}
