package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var filepath string
var verbose bool

type _instruction struct {
	count    int
	from, to string
}

func init() {
	flag.BoolVar(&verbose, "v", false, "show debug logs")
	flag.StringVar(&filepath, "f", "day6-file.input", "path to file")
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

func processLine(line string, lno int) {
	chars := make(map[rune][]int)
	for idx, b := range line {
		idxs, ok := chars[b]
		if !ok {
			idxs = []int{}
		}
		idxs = append(idxs, idx)
		chars[b] = idxs
		if idx > 3 {
			charToRemove := rune(line[idx-4])
			log.Printf("Removing char %v at %d", charToRemove, idx)
			found := chars[charToRemove]
			if len(found) > 1 {
				chars[charToRemove] = found[1:]
			} else {
				delete(chars, charToRemove)
			}
			log.Printf("line: %d, chars:%v", lno, chars)
		}

		if len(chars) == 4 {
			log.Printf("line: %d, found 4 keys at idx: %d (%v)", lno, idx, chars)
			fmt.Printf("line: %d, first loc of 4 non-repeating chars: %d\n", lno, idx+1)
			break
		}
	}
}

func main() {
	data := readFile(filepath)
	lines := strings.Split(string(data), "\n")
	for lno, line := range lines {
		log.Printf("line:%d, data:%s", lno, line)
		processLine(line, lno)
	}
}
