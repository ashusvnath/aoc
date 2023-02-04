package main

import (
	"day16/models"
	"day16/parser"
	"flag"
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
	lines := strings.Count(strings.TrimRight(string(data), "\n"), "\n") + 1
	if lines != len(models.AllRoomIds) {
		log.Fatalf("Something is very very wrong !!. %d, %d", lines, len(models.AllRoomIds))
	}
	floydWarshallPairDistances := FindAllPairPaths(models.AllRooms)
	//fmt.Printf("All pair distances: \n%v", floydWarshallPairDistances)
	//bfs := utility.NewBFS[string]()
	///Part1(bfs)
	Part1(floydWarshallPairDistances)
	// max, path := DFS(bfs, 0, 0, "AA", utility.NewSet[string](), []string{})
	// fmt.Printf("Max %d. Len: %d, Path: %v", max, len(path), path)
}
