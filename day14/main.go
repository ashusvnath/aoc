package main

import (
	"day14/models"
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
	Part1(grid)
}

func Part1(grid *models.Grid) {
	count := 0
	for fallenOff := false; !fallenOff; fallenOff = grid.Drop() {
		count++
		fmt.Printf("Grid:\n%s\n", grid)
	}
	fmt.Printf("Grid:\n%s\n", grid)
	fmt.Printf("Took %d steps\n", count)
}
