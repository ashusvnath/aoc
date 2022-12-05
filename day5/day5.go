package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

var filepath string
var verbose bool

func init() {
	flag.BoolVar(&verbose, "v", false, "show debug logs")
	flag.StringVar(&filepath, "f", "day5-file.input", "path to file")
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
	part1Result := ""
	part2Result := ""
	data := readFile(filepath)
	parts := strings.Split(string(data), "\n\n")
	currentStackString := parts[0]
	//instructions := parts[1]
	currentStackLines := strings.Split(currentStackString, "\n")
	currentStackLinesLen := len(currentStackLines)
	labelsRegexp := regexp.MustCompile(`\d+`)
	log.Printf("Looking for indexes in : %v", currentStackLines[currentStackLinesLen-1])
	indices := labelsRegexp.FindAllIndex([]byte(currentStackLines[currentStackLinesLen-1]), 100)
	log.Printf("Indices: %v", indices)

	//Initialize stacks
	stacks := make(map[string][]string, len(indices))
	stackIds := make([]string, len(indices))
	for stackId, idxs := range indices {
		stackIds[stackId] = currentStackLines[currentStackLinesLen-1][idxs[0]:idxs[1]]
		stacks[stackIds[stackId]] = make([]string, 0)
	}

	//Read each line of output in reverse order
	for i := currentStackLinesLen - 2; i >= 0; i-- {
		log.Printf("Reading: %v", currentStackLines[i])

		for iter, idxs := range indices {
			stackId := stackIds[iter]
			crateName := currentStackLines[i][idxs[0]:idxs[1]]
			if crateName != " " {
				stacks[stackId] = append(stacks[stackIds[iter]], crateName)
			}
		}
	}
	log.Printf("StackIds: %v\n", stackIds)
	log.Printf("Stacks: %v\n", stacks)

	//fmt.Printf("Current stack:\n%v\n\n", currentStackString)
	//fmt.Printf("Instructions:\n%v\n", instructions)
	fmt.Printf("Part 1: %s\n", part1Result)
	fmt.Printf("Part 2: %s\n", part2Result)
}
