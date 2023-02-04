package main

import "day16/models"

type Pair [2]string

func FindAllPairPaths(allRooms map[string]*models.Room) map[Pair]int {
	pairDistances := make(map[Pair]int)
	for id, room := range allRooms {
		pairDistances[Pair{id, id}] = 0
		for _, n := range models.AllRoomIds {
			if room.ConnectedTo(n) {
				pairDistances[Pair{id, n}] = 1
			} else {
				pairDistances[Pair{id, n}] = len(models.AllRoomIds) + 1
			}
		}
	}
	for _, between := range models.AllRoomIds {
		for _, from := range models.AllRoomIds {
			if between == from {
				continue
			}
			for _, to := range models.AllRoomIds {
				if from == to || to == between {
					continue
				}
				pdft := pairDistances[Pair{from, to}]
				pdfb := pairDistances[Pair{from, between}]
				pdbt := pairDistances[Pair{between, to}]
				if pdft > pdfb+pdbt {
					pairDistances[Pair{from, to}] = pdfb + pdbt
				}
			}
		}
	}
	return pairDistances
}
