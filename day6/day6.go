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
var cpuprofile *string

type _instruction struct {
	count    int
	from, to string
}

func init() {
	flag.BoolVar(&verbose, "v", false, "show debug logs")
	flag.StringVar(&filepath, "f", "day6-file.input", "path to file")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profiling info to file")
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

func processLineForSOP(line string, lno int) {
	sopChars := make(map[rune][]int)
	for idx, b := range line {

		idxs, ok := sopChars[b]
		if !ok {
			idxs = []int{}
		}
		idxs = append(idxs, idx)
		sopChars[b] = idxs

		if idx > 3 {
			charToRemove := rune(line[idx-4])
			log.Printf("Removing char %v at %d", charToRemove, idx)
			foundSop := sopChars[charToRemove]
			if len(foundSop) > 1 {
				sopChars[charToRemove] = foundSop[1:]
			} else {
				delete(sopChars, charToRemove)
			}
		}

		log.Printf("line: %d, SOP chars:%v", lno, sopChars)

		if len(sopChars) == 4 {
			log.Printf("line: %d, found 4 keys at idx: %d (%v)", lno, idx, sopChars)
			fmt.Printf("line: %d, first loc of 4 non-repeating chars: %d\n", lno, idx+1)
			break
		}
	}
}

func processLineForSOM(line string, lno int) {
	somChars := make(map[rune][]int)
	for idx, b := range line {
		idxs, ok := somChars[b]
		if !ok {
			idxs = []int{}
		}
		idxs = append(idxs, idx)
		somChars[b] = idxs
		if idx > 13 {
			charToRemove := rune(line[idx-14])
			log.Printf("Removing char %v at %d", charToRemove, idx)
			foundSom := somChars[charToRemove]
			if len(foundSom) > 1 {
				somChars[charToRemove] = foundSom[1:]
			} else {
				delete(somChars, charToRemove)
			}
		}
		log.Printf("line: %d, SOM chars: %v", lno, somChars)

		if len(somChars) == 14 {
			log.Printf("line: %d, found 14 keys at idx: %d (%v)", lno, idx, somChars)
			fmt.Printf("line: %d, first loc of 14 non-repeating chars: %d\n", lno, idx+1)
			break
		}

	}
}

func main() {
	data := readFile(filepath)
	lines := strings.Split(string(data), "\n")
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	for sampleCount := 0; sampleCount < 1000; sampleCount++ {
		for lno, line := range lines {
			log.Printf("line:%d, data:%s", lno, line)
			processLineForSOP(line, lno)
			processLineForSOM(line, lno)
		}
	}

}
