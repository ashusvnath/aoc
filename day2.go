package main

import (
	"os"
	"fmt"
	"flag"
	"log"
	"strings"
)

var filepath string
var verbose bool
var equivalents, beatenBy map[string]string
var scoring map[string]int


func init() {
	beatenBy = make(map[string]string, 3)
	beatenBy["A"] = "Y";	beatenBy["B"] = "Z"; beatenBy["C"] = "X" 

	equivalents = make(map[string]string, 3)
	equivalents["A"] = "X";	equivalents["B"] = "Y"; equivalents["C"] = "Z" 

	scoring = make(map[string]int, 3)
	scoring["X"] = 1; scoring["Y"] = 2; scoring["Z"] = 3;

	flag.BoolVar(&verbose, "v", false, "show debug logs")
	flag.StringVar(&filepath, "f", "day2-input.txt", "path to file")
	flag.Parse()
}

func main() {
	data, err := os.ReadFile(filepath)
    if err != nil {
		log.Printf("Could not read contents of %s : %v", filepath, err)
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")
	score := 0
	totalScore := 0

	for _, line := range lines {
		game := strings.Split(line, " ")
		score = scoring[game[1]]
		if beatenBy[game[0]] == game[1] {
			score += 6
		} else if equivalents[game[0]] == game[1] {
			score += 3
		}
		totalScore += score
		log.Printf("line: %s, score: %d, totalScore: %d", line, score, totalScore)
	}
	fmt.Printf("Part 1 : %d\n", totalScore)
}