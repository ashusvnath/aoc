package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"math/cmplx"
	"os"
	"regexp"
	"runtime/pprof"
	"strconv"
	"strings"
)

var filepath string
var verbose bool
var cpuprofile string

var moveDelta = map[string]complex128{
	"R": 1,
	"L": -1,
	"U": 1i,
	"D": -1i,
}

type Rope struct {
	head          complex128
	tail          complex128
	tailPositions Set[complex128]
}

func (r *Rope) Move(direction string, count int) {
	for i := 0; i < count; i++ {
		r.head += moveDelta[direction]
		delta := r.head - r.tail
		l, _ := cmplx.Polar(delta)
		if l >= 2 {
			r.moveTail(delta)
		}
		log.Printf("Move: %s Rope: %#v", direction, r)
	}
}

func (r *Rope) moveTail(delta complex128) {
	dx, dy := real(delta), imag(delta)
	if dx != 0 {
		dx = dx / math.Abs(dx)
	}
	if dy != 0 {
		dy = dy / math.Abs(dy)
	}

	r.tail += complex(dx, dy)
	r.tailPositions.Add(r.tail)
}

func (r *Rope) CountDistinctTailPositions() int {
	return len(r.tailPositions)
}

func NewRope() *Rope {
	positions := make(Set[complex128])
	positions.Add(0)
	return &Rope{0, 0, positions}
}

func init() {
	flag.BoolVar(&verbose, "v", false, "show debug logs")
	flag.StringVar(&filepath, "f", "test.txt", "path to file")
	flag.StringVar(&cpuprofile, "cpuprofile", "", "write cpu profile to file")
	flag.Parse()
	if !verbose {
		log.SetOutput(io.Discard)
	}
}

func readFile(filepath string) []byte {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Printf("Could not read contents of %s : %v", filepath, err)
		os.Exit(1)
	}
	return data
}

func main() {
	data := readFile(filepath)
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	log.Printf("Input: \n%v\n", string(data))
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	rope := NewRope()
	commandRegexp := regexp.MustCompile(`^([RULD]) (\d+)`)
	for _, line := range lines {
		result := commandRegexp.FindSubmatch([]byte(line))
		if result == nil {
			continue
		}
		direction := string(result[1])
		count, _ := strconv.Atoi(string(result[2]))
		log.Printf("Direction: %s, count: %v", direction, count)
		rope.Move(direction, count)
	}
	fmt.Printf("Part1: %d\n", rope.CountDistinctTailPositions())
}
