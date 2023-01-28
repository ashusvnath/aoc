package models

import (
	"bytes"
	"log"
	"math"
	"strconv"
	"strings"
)

type Location complex128

const None Location = complex(-1, -1)
const Source Location = complex(500, 0)

func (l Location) Next() []Location {
	return []Location{
		l + 1i,     // down
		l - 1 + 1i, // left
		l + 1 + 1i, // right
	}
}

type Kind int

const (
	EMPTY Kind = 0
	ROCK  Kind = 1
	SAND  Kind = 2
)

type Object struct {
	kind     Kind
	location Location
}

type Grid struct {
	rawData       string
	mat           map[Location]Object
	left, right   float64
	floor, bottom float64
}

func (g *Grid) RawData() string {
	return g.rawData
}

func (g *Grid) Occupied(location Location) bool {
	_, found := g.mat[location]
	return found
}

func (g *Grid) NextValid(currentLocation Location) Location {
	for _, l := range currentLocation.Next() {
		if !g.Occupied(l) && imag(l) <= g.floor {
			return l
		}
	}
	return None
}

func ParseGrid(data string) *Grid {
	input := strings.TrimSuffix(string(data), "\n")
	matrix := make(map[Location]Object)
	grid := &Grid{
		rawData: data,
		mat:     matrix,
		left:    100000,
		right:   0,
		bottom:  -1,
		floor:   100000,
	}
	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		vertices := []Location{}
		for _, vertexCoordinates := range strings.Split(line, " -> ") {
			vertexCoordinateStrings := strings.Split(vertexCoordinates, ",")
			x, _ := strconv.ParseFloat(vertexCoordinateStrings[0], 64)
			y, _ := strconv.ParseFloat(vertexCoordinateStrings[1], 64)
			vertices = append(vertices, Location(complex(x, y)))
			grid.fill(vertices)
		}
	}
	log.Printf("left: %v, right:%v, bottom: %v", grid.left, grid.right, grid.bottom)
	grid.floor = grid.bottom + 1
	return grid
}

func (grid *Grid) Add(o Object) {
	point := o.location
	grid.mat[o.location] = o
	x, y := real(point), imag(point)
	if grid.left > x {
		grid.left = x
	} else if grid.right < x {
		grid.right = x
	}
	if grid.floor > y && grid.bottom < y {
		grid.bottom = y
	}
}

func (grid *Grid) fill(vertices []Location) {
	for idx, vertex := range vertices[:len(vertices)-1] {
		start := vertex
		end := vertices[idx+1]
		delta := end - start
		absDelta := math.Abs(real(delta) + imag(delta))
		delta = Location(complex(real(delta)/absDelta, imag(delta)/absDelta))
		for point := start; point != end; point += delta {
			grid.Add(Object{ROCK, point})
		}
		grid.Add(Object{ROCK, end})
	}
}

func (grid *Grid) String() string {
	helper := map[Kind]byte{
		ROCK:  '#',
		SAND:  'o',
		EMPTY: '.',
	}
	output := &bytes.Buffer{}
	for y := float64(0); y <= float64(grid.floor)+1; y += 1 {
		for x := float64(grid.left) - 2; x < float64(grid.right)+2; x += 1 {
			point := Location(complex(x, y))
			kind := grid.mat[point].kind
			if y == grid.floor+1 {
				kind = ROCK
			}
			output.WriteByte(helper[kind])
		}
		output.WriteByte('\n')
	}
	return output.String()
}

func (g *Grid) WithinBounds(location Location) bool {
	return g.left <= real(location) && g.right >= real(location) && g.bottom >= imag(location)
}

func (grid *Grid) Drop() bool {
	start := Location(Source)
	fallenOff := false
	for grid.NextValid(start) != None {
		start = grid.NextValid(start)
		if !grid.WithinBounds(start) {
			fallenOff = true
			break
		}
	}
	if !fallenOff {
		grid.Add(Object{SAND, start})
	}
	return fallenOff
}

func (grid *Grid) OnFloor(location Location) bool {
	return imag(location) == grid.floor
}

func (grid *Grid) DropToFloor() bool {
	point := Source
	noValidNext := true
	for point = grid.NextValid(point); point != None; {
		noValidNext = false
		if grid.OnFloor(point) {
			grid.Add(Object{SAND, point})
			break
		}
		next := grid.NextValid(point)
		if next == None {
			grid.Add(Object{SAND, point})
		}
		point = next
	}
	if noValidNext {
		grid.Add(Object{SAND, Source})
	}
	return noValidNext
}
