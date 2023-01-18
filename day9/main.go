package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"runtime/pprof"
	"strconv"
	"strings"
)

var filepath string
var verbose bool
var cpuprofile string
var commandRegexp *regexp.Regexp

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

func init() {
	commandRegexp = regexp.MustCompile(`^([RULD]) (\d+)`)
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
	rope := NewRope(10)
	node2Positions := NewRecorder()
	node10Positions := NewRecorder()

	rope.RegisterRecorderByKnotIdx(node2Positions.Record, 2)
	rope.RegisterRecorderByKnotIdx(node10Positions.Record, 10)
	log.Printf("Position: %#v\n", rope)
	Execute(lines, rope)

	fmt.Printf("Part1: %d\n", node2Positions.Count())
	fmt.Printf("Part2: %d\n", node10Positions.Count())
}

func Execute(lines []string, rope *Rope) {
	for _, line := range lines {
		result := commandRegexp.FindSubmatch([]byte(line))
		if result == nil {
			continue
		}
		direction := string(result[1])
		count, _ := strconv.Atoi(string(result[2]))
		log.Printf("Execute: Direction: %s, count: %v", direction, count)
		rope.Move(direction, count)
	}
}
