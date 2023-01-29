package main

import (
	"day15/models"
	"day15/parser"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
)

var filepath string
var verbose bool
var cpuprofile string

func init() {
	flag.BoolVar(&verbose, "v", false, "show debug logs")
	flag.StringVar(&filepath, "f", "test.txt", "path to file")
	flag.StringVar(&cpuprofile, "cpuprofile", "", "write cpu profile to file")
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
	var targetY int
	flag.IntVar(&targetY, "t", 10, "target line y=?")
	flag.Parse()
	if !verbose {
		log.SetOutput(io.Discard)
	}

	data := readFile(filepath)
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	part1 := models.NewSensorProximityObserver(float64(targetY))
	parser.Parse(string(data), part1.Listen)
	log.Printf("Data:\n%v", string(data))
	fmt.Printf("Part1: %d\n", part1.Count())
}
