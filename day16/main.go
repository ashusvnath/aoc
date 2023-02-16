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
	bitMask := utility.NewBitmask(models.AllRoomIds)
	pf := NewPathFinder(models.AllRooms, floydWarshallPairDistances, bitMask)
	prl := NewParialResultListener(floydWarshallPairDistances, 26, bitMask)
	pf.SetListener(prl.Listen)
	visited := utility.NewBitMaskSet(bitMask)
	pf.SetTimeLimit(30)
	result, _, err := pf.Search("AA", visited, 30, 0, []string{"AA"}, "man")

	if err != nil {
		log.Fatalf("Part1: Error encountered when searching for best path %v", result)
	}
	fmt.Printf("AllRoomUsefulIds: %v, %d\n", models.RoomIDsWithNonZeroReleaseRate, len(models.RoomIDsWithNonZeroReleaseRate))
	fmt.Printf("Part1: %d, %v, %v\n", result.totalReleased, result.opened, result)
	//path, _, _ := RenderPath(result.opened, pf, 30)
	//fmt.Printf("%s\n\n", path)

	fmt.Println("Part2:")
	results, total := prl.BestResult()
	fmt.Printf("Total: %d,\n Result1: %v,\n Result2: %v\n", total, results[0], results[1])
}
