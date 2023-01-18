package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	monkeyLineRegex = regexp.MustCompile(`^Monkey (\d+):$`)
	itemsLineRegex  = regexp.MustCompile(`^\s+Starting items: (\d+(,\s\d+)*)$`)
	operationRegex  = regexp.MustCompile(`^\s+Operation: (new = old ([+*]) (old|\d+))$`)
)

type Parser struct {
	currentMonkeyBuilder *MonkeyBuilder
}

func (p *Parser) Parse(monkeyLines string) {
	lines := strings.Split(monkeyLines, "\n")
	p.currentMonkeyBuilder = NewMonkeyBuilder()
	for _, line := range lines {
		switch {
		case monkeyLineRegex.MatchString(line):
			id := monkeyLineRegex.FindStringSubmatch(line)[1]
			p.currentMonkeyBuilder.Id(id)
		case itemsLineRegex.MatchString(line):
			itemsString := itemsLineRegex.FindStringSubmatch(line)[1]
			itemsAsStrings := strings.Split(itemsString, ", ")
			items := make([]int, len(itemsAsStrings))
			for i, item := range itemsAsStrings {
				items[i], _ = strconv.Atoi(item)
			}
			p.currentMonkeyBuilder.Items(items...)
		case operationRegex.MatchString(line):
			operationStrings := operationRegex.FindStringSubmatch(line)
			operation := ParseOperation(operationStrings)
			p.currentMonkeyBuilder.Op(operation)
		}

	}
}

func ParseOperation(input []string) Operation {
	operation := input[2]
	operand := input[3]
	switch operation {
	case "*":
		if operand == "old" {
			return Square()
		}
		opInt, _ := strconv.Atoi(operand)
		return Mul(opInt)
	case "+":
		opInt, _ := strconv.Atoi(operand)
		return Add(opInt)
	default:
		log.Fatalf("Unknown operation %s", input[1])
	}
	return nil
}
