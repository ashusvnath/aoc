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
	log.Printf("Data: \n%v\n", string(data))
	cpu := NewCPU(NewClock())
	part1 := NewRecorder("part1", 20, 40)
	part2 := NewScreen(6, 40, '#', '.')
	cpu.Register(part1)
	cpu.Register(part2)
	cpu.Execute(strings.Split(strings.TrimRight(string(data), "\n"), "\n"))
	fmt.Printf("Part1: %v\n", part1.Report())
	fmt.Printf("Part2: \n%v", part2.Report())
}
