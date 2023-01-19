package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
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

	log.Printf("Data: %v", string(data))
	fmt.Printf("Part1: %v\n", Part1(string(data)))
	fmt.Printf("Part2: %v\n", Part2(string(data)))
}

func Part1(data string) int {
	ms := NewMonkeyService(Divide(big.NewInt(3)))
	ms.Setup(string(data))
	ms.DoBusiness(20)
	monkeys := ms.MonkeysByActivity()
	log.Printf("Part1: sorted by activity : %v", monkeys)

	result := monkeys[0].Activity() * monkeys[1].Activity()
	return result
}

func Part2(data string) int {
	ms := NewMonkeyService(nil)
	ms.Setup(string(data))
	ms.DoBusiness(10000)
	monkeys := ms.MonkeysByActivity()
	log.Printf("Part2: sorted by activity : %v", monkeys)

	result := monkeys[0].Activity() * monkeys[1].Activity()
	return result
}
