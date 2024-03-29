package main

import (
	"day13/models"
	"day13/parser"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"sort"
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
	pairs, allPackets := parser.ParsePairs(string(data))
	fmt.Printf("Part1: %d\n", Part1(pairs))
	fmt.Printf("Part2: %d\n", Part2(allPackets))

}

func Part1(pairs []*models.Pair) int {
	sum := 0
	for idx, pair := range pairs {
		log.Printf("=== Pair %d ===", idx+1)
		if pair.IsOrderedCorrectly() {
			log.Printf("Pair at idx %d is ordered correctly", idx+1)
			sum += (idx + 1)
		}
	}
	return sum
}

func Part2(packets []models.List) int {
	divider2 := models.List{models.List{models.Int(2)}}
	divider6 := models.List{models.List{models.Int(6)}}
	packets = append(packets, divider2, divider6)
	sort.Slice(packets, func(left, right int) bool {
		return packets[left].Compare(packets[right]) == models.Less
	})
	result := 1
	for idx, p := range packets {
		if p.Compare(divider2) == models.Equal || p.Compare(divider6) == models.Equal {
			result *= (idx + 1)
		}
	}
	return result
}
