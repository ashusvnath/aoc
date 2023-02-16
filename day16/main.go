package main

import (
	"day16/models"
	"day16/parser"
	"day16/utility"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"strings"
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
	parser.Parse(strings.TrimRight(string(data), "\n"))
	log.Printf("AllRooms:\n %v", models.AllRooms)

	floydWarshallPairDistances := FindAllPairPaths(models.AllRooms)
	pf := NewPathFinder(models.AllRooms, floydWarshallPairDistances)
	Part1(pf)
	Part2(pf)
}

func Part1(pf *PathFinder) {
	visited := utility.NewSet[string]()
	pf.SetTimeLimit(30)
	result, _, err := pf.Search("AA", visited, 30, 0, []string{"AA"}, "man")
	if err != nil {
		log.Fatalf("Part1: Error encountered when searching for best path %v", result)
	}
	fmt.Printf("Part1: %d, %v\n", result.totalReleased, result.opened)
	path, _, _ := RenderPath(result.opened, pf, 30)
	fmt.Printf("%s\n\n", path)
}

func Part2(pf *PathFinder) {
	pf.SetTimeLimit(26)
	result, err := pf.SearchWithCache("AA", 26, "elephant")
	if err != nil {
		log.Fatalf("Error when searching for best path with cachedResults: %v", err)
	}

	path1, released1, timeTaken1 := RenderPath(result[0].opened, pf, 26)
	path2, released2, timeTaken2 := RenderPath(result[1].opened, pf, 26)
	totalReleased := released1 + released2
	timeTaken := -1
	if timeTaken1 > timeTaken2 {
		timeTaken = timeTaken1
	} else {
		timeTaken = timeTaken2
	}
	fmt.Printf("Part2: %v, %v\n", result[0], result[1])
	fmt.Printf("%s - \n%s\nReleased: %d\nTotalTime taken: %d\n", path1, path2, totalReleased, timeTaken)
}
