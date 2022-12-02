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
var draws, beatenBy, winsAgainst map[string]string
var scoring map[string]int


func init() {
	beatenBy = make(map[string]string, 3)
	beatenBy["A"] = "Y";	beatenBy["B"] = "Z"; beatenBy["C"] = "X" 

	draws = make(map[string]string, 3)
	draws["A"] = "X";	draws["B"] = "Y"; draws["C"] = "Z" 

	scoring = make(map[string]int, 6)
	scoring["X"] = 1; scoring["Y"] = 2; scoring["Z"] = 3;
	scoring["A"] = 1; scoring["B"] = 2; scoring["B"] = 3;

	winsAgainst = make(map[string]string, 3)
	winsAgainst["A"] = "Z";	winsAgainst["B"] = "X"; winsAgainst["C"] = "Y" 

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
	part1score := 0
	part1totalScore := 0
	part2score := 0
	part2totalScore := 0

	for _, line := range lines {
		game := strings.Split(line, " ")
		part1score = scoring[game[1]]
		if beatenBy[game[0]] == game[1] {
			part1score += 6
		} else if draws[game[0]] == game[1] {
			part1score += 3
		}
		part1totalScore += part1score
		log.Printf("Part1 -> line: %s, score: %d, totalScore: %d", line, part1score, part1totalScore)

		switch(game[1]) {
		case "X":
			part2score = 0 + scoring[winsAgainst[game[0]]]
		case "Y":
			part2score = 3 + scoring[draws[game[0]]]
		case "Z":
			part2score = 6 + scoring[beatenBy[game[0]]]
		}
		part2totalScore += part2score
		log.Printf("Part2 -> line: %s, score: %d, totalScore: %d", line, part2score, part2totalScore)
	}
	fmt.Printf("Part 1 : %d\n", part1totalScore)
	fmt.Printf("Part 2 : %d\n", part2totalScore)
}