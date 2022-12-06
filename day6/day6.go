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

type UniquePrefixDetector struct {
	chars     map[rune][]int
	prefixLen int
	lastIndex int
	seen      string
	Found     bool
	AddRune   func(u *UniquePrefixDetector, b rune, idx int)
}

func NewUniquePrefixDetector(length int) *UniquePrefixDetector {
	return &UniquePrefixDetector{
		chars:     make(map[rune][]int),
		prefixLen: length,
		lastIndex: -1,
		seen:      "",
		Found:     false,
		AddRune:   addRune,
	}
}

func addRune(u *UniquePrefixDetector, b rune, idx int) {
	u.seen += string(b)
	idxs, ok := u.chars[b]
	if !ok {
		idxs = []int{}
	}
	idxs = append(idxs, idx)
	u.chars[b] = idxs
	u.lastIndex = idx
	processDuplicates(u)
}

func processDuplicates(u *UniquePrefixDetector) {
	idx := u.lastIndex
	if idx > u.prefixLen-1 {
		charToRemove := rune(u.seen[idx-u.prefixLen])
		log.Printf("Removing char %v at %d", charToRemove, idx)
		foundIdxs := u.chars[charToRemove]
		if len(foundIdxs) > 1 {
			u.chars[charToRemove] = foundIdxs[1:]
		} else {
			delete(u.chars, charToRemove)
		}
	}
	u.Found = len(u.chars) == u.prefixLen
	if u.Found {
		u.AddRune = NOPaddRune
	}
}

func NOPaddRune(_ *UniquePrefixDetector, _ rune, _ int) {
	return
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

func main() {
	data := readFile(filepath)
	lines := strings.Split(string(data), "\n")
	for lno, line := range lines {
		prefix4 := NewUniquePrefixDetector(4)
		prefix14 := NewUniquePrefixDetector(14)
		for idx, b := range line {
			prefix4.AddRune(prefix4, b, idx)
			prefix14.AddRune(prefix14, b, idx)
			if prefix4.Found && prefix14.Found {
				break
			}
		}
		fmt.Printf("line: %d, first loc of 4 non-repeating chars: %d\n", lno, prefix4.lastIndex+1)
		fmt.Printf("line: %d, first loc of 14 non-repeating chars: %d\n", lno, prefix14.lastIndex+1)
		log.Printf("line:%d, data:%s", lno, line)
	}
}
