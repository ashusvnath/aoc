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

var filepath string
var verbose bool

type _instruction struct {
	count    int
	from, to string
}

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

func applyInstructionPart1(inst _instruction, stacks map[string][]string) {
	//read
	from := stacks[inst.from]
	to := stacks[inst.to]

	//transform
	stackLen := len(from)
	taken := from[stackLen-inst.count:]
	from = from[0 : stackLen-inst.count]
	for idx := len(taken) - 1; idx >= 0; idx-- {
		to = append(to, taken[idx])
	}

	//store
	stacks[inst.from] = from
	stacks[inst.to] = to
}

func applyInstructionPart2(inst _instruction, stacks map[string][]string) {
	//read
	from := stacks[inst.from]
	to := stacks[inst.to]

	//transform
	stackLen := len(from)
	taken := from[stackLen-inst.count:]
	from = from[0 : stackLen-inst.count]
	to = append(to, taken...)

	//store
	stacks[inst.from] = from
	stacks[inst.to] = to
}

func readStack(currentStackString string) ([2]map[string][]string, []string) {
	currentStackLines := strings.Split(currentStackString, "\n")
	currentStackLinesLen := len(currentStackLines)
	labelsRegexp := regexp.MustCompile(`\d+`)
	log.Printf("Looking for indexes in : %v", currentStackLines[currentStackLinesLen-1])
	indices := labelsRegexp.FindAllIndex([]byte(currentStackLines[currentStackLinesLen-1]), 100)
	log.Printf("Indices: %v", indices)

	//Initialize stacks
	stacks0 := make(map[string][]string, len(indices))
	stacks1 := make(map[string][]string, len(indices))
	stackIds := make([]string, len(indices))
	for stackId, idxs := range indices {
		stackIds[stackId] = currentStackLines[currentStackLinesLen-1][idxs[0]:idxs[1]]
		stacks0[stackIds[stackId]] = make([]string, 0)
		stacks1[stackIds[stackId]] = make([]string, 0)
	}

	//Read each line of output in reverse order
	for i := currentStackLinesLen - 2; i >= 0; i-- {
		log.Printf("Reading: %v", currentStackLines[i])

		for iter, idxs := range indices {
			stackId := stackIds[iter]
			crateName := currentStackLines[i][idxs[0]:idxs[1]]
			if crateName != " " {
				stacks0[stackId] = append(stacks0[stackIds[iter]], crateName)
				stacks1[stackId] = append(stacks1[stackIds[iter]], crateName)
			}
		}
	}
	return [2]map[string][]string{stacks0, stacks1}, stackIds
}

func processInstructions(input string, stacks [2]map[string][]string, stackIds []string) {
	instuructionPattern := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	instructions := strings.Split(input, "\n")
	for _, instruction := range instructions {
		log.Printf("Parsing Instruction: %v", instruction)
		matches := instuructionPattern.FindSubmatch([]byte(instruction))
		count, _ := strconv.Atoi(string(matches[1]))
		from := string(matches[2])
		to := string(matches[3])
		log.Printf("Instruction: move count:%d, from:%s , to:%s", count, from, to)
		applyInstructionPart1(_instruction{count, from, to}, stacks[0])
		applyInstructionPart2(_instruction{count, from, to}, stacks[1])
		log.Printf("After apply: %v", stacks)
	}
}

func main() {
	part1Result := ""
	part2Result := ""
	data := readFile(filepath)
	parts := strings.Split(string(data), "\n\n")
	stacks, stackIds := readStack(parts[0])

	log.Printf("StackIds: %v\n", stackIds)
	log.Printf("Stacks: %v\n", stacks)

	processInstructions(parts[1], stacks, stackIds)
	for _, stackId := range stackIds {
		part1Result += stacks[0][stackId][len(stacks[0][stackId])-1]
		part2Result += stacks[1][stackId][len(stacks[1][stackId])-1]
	}

	log.Printf("Stacks : %v", stacks)
	fmt.Printf("Part 1: %s\n", part1Result)
	fmt.Printf("Part 2: %s\n", part2Result)
}
