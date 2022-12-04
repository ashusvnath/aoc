package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type _range struct {
	start, end int
}

func (this *_range) contains(other *_range) bool {
	return this.start <= other.start && this.end >= other.end
}

var filepath string
var verbose bool

func init() {
	flag.BoolVar(&verbose, "v", false, "show debug logs")
	flag.StringVar(&filepath, "f", "day4-file.input", "path to file")
	flag.Parse()
	if !verbose {
		log.SetOutput(io.Discard)
	}
}

func main() {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Printf("Could not read contents of %s : %v", filepath, err)
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")

	overlapping := 0
	rangeExp := regexp.MustCompile(`(?P<start1>\d+)-(?P<end1>\d+),(?P<start2>\d+)-(?P<end2>\d+)`)
	var matches [][]byte
	for idx, line := range lines {
		matches = rangeExp.FindSubmatch([]byte(line))
		log.Printf("%v", matches)
		if len(matches) == 0 {
			continue
		}
		leftStart, _ := strconv.Atoi(string(matches[1]))
		leftEnd, _ := strconv.Atoi(string(matches[2]))
		leftRange := &_range{leftStart, leftEnd}

		rightStart, _ := strconv.Atoi(string(matches[3]))
		rightEnd, _ := strconv.Atoi(string(matches[4]))
		rightRange := &_range{rightStart, rightEnd}
		log.Printf("line(%d): %s ; Ranges: %v, %v", idx, line, leftRange, rightRange)

		if leftRange.contains(rightRange) || rightRange.contains(leftRange) {
			log.Printf("Found an overlap: %v, %v", leftRange, rightRange)
			overlapping += 1
		}
	}
	fmt.Printf("Part 1: %d\n", overlapping)
}
