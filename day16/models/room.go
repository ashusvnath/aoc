package models

import (
	"day16/utility"
	"fmt"
)

type Room struct {
	id                  string
	releaseRate         int
	connectedRoomIdsSet *utility.Set[string]
	valveOpened         bool
	__str               string
}

func (r *Room) GoString() string {
	return r.__str
}

func (r *Room) Id() string {
	return r.id
}

func (r *Room) ReleaseRate() int {
	return r.releaseRate
}

func (r *Room) ConnectedRoomIds() []string {
	return r.connectedRoomIdsSet.AsSlice()
}

func (r *Room) IsValveOpened() bool {
	return r.valveOpened
}

func (r *Room) ConnectedTo(otherRoomId string) bool {
	return r.connectedRoomIdsSet.Contains(otherRoomId)
}

var AllRooms = make(map[string]*Room)
var AllRoomIds = make([]string, 0)
var RoomIDsWithNonZeroReleaseRate = []string{}

func NewRoom(id string, rate int, connectedRoomIds ...string) *Room {
	connectedRoomIdsSet := utility.NewSet[string]()
	for _, id := range connectedRoomIds {
		connectedRoomIdsSet.Add(id)
	}

	AllRoomIds = append(AllRoomIds, id)
	AllRooms[id] = &Room{
		id:                  id,
		releaseRate:         rate,
		connectedRoomIdsSet: connectedRoomIdsSet,
		valveOpened:         false,
		__str:               fmt.Sprintf("Room{Id:%s, Neighbours:(%v)}", id, connectedRoomIds),
	}
	if rate > 0 {
		RoomIDsWithNonZeroReleaseRate = append(RoomIDsWithNonZeroReleaseRate, id)
	}
	return AllRooms[id]
}
