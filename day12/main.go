package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"

	"day12/models"
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
	g := models.ParseGrid(data)
	pf := NewPathFinder(g)
	path := pf.FindPath(g.Start())
	rand.Seed(time.Now().UnixNano())
	log.Printf("Path visualized:\n%s", VisualizePath(g, path.GetLocations(), path.Start()))
	fmt.Printf("Part1:length:%d\n", path.Len())
	trail := pf.FindHikingTrail(path)
	log.Printf("Searched %d starting points", g.GetLocationsAtHeight(0).Len())
	log.Printf("Shortest trail visualized:\n%s", VisualizePath(g, trail.GetLocations(), trail.Start()))
	fmt.Printf("Part2:length:%d\n", trail.Len())
}
