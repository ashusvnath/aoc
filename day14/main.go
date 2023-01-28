package main

import (
	"day14/models"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var filepath string
var verbose bool
var cpuprofile string
var renderIntermediates bool

func init() {
	flag.BoolVar(&verbose, "v", false, "show debug logs")
	flag.BoolVar(&renderIntermediates, "r", false, "show debug logs")
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
	log.Printf("Data:\n%v", string(data))
	grid := models.ParseGrid(string(data))
	count := Part1(grid)
	fmt.Printf("Part1: %d\n", count)
	count = Part2(grid, count)
	fmt.Printf("Part2: %d\n", count)

}

func Part1(grid *models.Grid) int {
	count := 0
	for fallenOff := false; !fallenOff; fallenOff = grid.Drop() {
		count++
		if renderIntermediates {
			fmt.Printf("Grid:\n%s\n", grid)
			time.Sleep(200 * time.Millisecond)
		}
	}
	fmt.Printf("Grid:\n%s\n", grid)
	return count - 1
}

func Part2(grid *models.Grid, count int) int {
	for stopped := false; !stopped; stopped = grid.DropToFloor() {
		count++
		if renderIntermediates {
			fmt.Printf("Grid:\n%s\n", grid)
			time.Sleep(200 * time.Millisecond)
		}
	}
	fmt.Printf("Grid:\n%s\n", grid)
	return count
}
