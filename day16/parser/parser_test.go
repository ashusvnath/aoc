package parser

import (
	"day16/assert"
	"day16/models"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("should parse single line and identify, flow rate and connected rooms", func(t *testing.T) {
		input := `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB`

		room := Parse(input)

		assert.True(room != nil, t)
		assert.Equal("AA", room.Id(), t)
		assert.Equal(0, room.ReleaseRate(), t)
		assert.False(room.IsValveOpened(), t)
		for _, connectedRoomId := range []string{"DD", "II", "BB"} {
			assert.True(room.ConnectedTo(connectedRoomId), t)
		}
	})

	t.Run("should parse multiple lines and identify, flow rate and connected rooms for each", func(t *testing.T) {
		input := `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnel leads to valve CC
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE`

		Parse(input)

		assert.Equal(3, len(models.AllRooms), t)
		assert.Equal(0, models.AllRooms["AA"].ReleaseRate(), t)
		assert.Equal(13, models.AllRooms["BB"].ReleaseRate(), t)
		assert.Equal(20, models.AllRooms["DD"].ReleaseRate(), t)
	})

}
