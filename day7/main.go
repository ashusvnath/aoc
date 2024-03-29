package main

import (
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
	flag.StringVar(&filepath, "f", "day6-file.input", "path to file")
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

	lines := strings.Split(string(data), "\n")

	rootDir := NewDir("/", nil)
	parser := NewParser(rootDir)
	for _, line := range lines {
		parser.Parse(line)
	}

	fmt.Printf("Known dirs: %d\n", len(knownDirs))
	GenerateListing(rootDir, 0)
	fmt.Printf("Part1: %d\n", Part1(knownDirs))
	fmt.Printf("Part2: %d\n", Part2(rootDir, knownDirs))
}

func Part1(knownDirs []*AOCDir) int {
	limit := 100000
	total := 0
	for _, dir := range knownDirs {
		size := dir.Size()
		if size <= limit {
			total += size
			log.Printf("Dir: %#v, Size : %#v, Running total: %#v\n", dir.name, size, total)
		}
	}
	return total
}

func Part2(rootDir *AOCDir, knownDirs []*AOCDir) int {
	totalDiskSize := 70000000
	requiredSpace := 30000000
	currentFree := totalDiskSize - rootDir.Size()

	minCleanupSize := requiredSpace - currentFree
	var smallestEligible *AOCDir = rootDir
	for _, dir := range knownDirs {
		if dir.Size() >= minCleanupSize && smallestEligible.Size() > dir.Size() {
			smallestEligible = dir
		}
	}
	return smallestEligible.size
}
