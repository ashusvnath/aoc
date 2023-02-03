package models

import "fmt"

type Room struct {
	id               string
	releaseRate      int
	connectedRoomIds []string
	valveOpened      bool
}

func (r *Room) GoString() string {
	return fmt.Sprintf("Room{Id:%s: Neighbours:(%v)}", r.id, r.connectedRoomIds)
}

func (r *Room) Id() string {
	return r.id
}

func (r *Room) ReleaseRate() int {
	return r.releaseRate
}

func (r *Room) ConnectedRoomIds() []string {
	return r.connectedRoomIds
}

func (r *Room) IsValveOpened() bool {
	return r.valveOpened
}

var AllRooms = make(map[string]*Room)
var RoomsWithNonZeroReleaseRate = 0

func NewRoom(id string, rate int, connectedRoomIds ...string) *Room {
	AllRooms[id] = &Room{
		id:               id,
		releaseRate:      rate,
		connectedRoomIds: connectedRoomIds,
		valveOpened:      false,
	}
	if rate > 0 {
		RoomsWithNonZeroReleaseRate++
	}
	return AllRooms[id]
}
