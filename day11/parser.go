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
	testLineRegex   = regexp.MustCompile(`^\s+Test: divisible by (\d+)$`)
	actionLineRegex = regexp.MustCompile(`^\s+If (true|false): throw to monkey (\d+)$`)
)

type Parser struct {
	currentMonkeyBuilder *MonkeyBuilder
}

func NewParser() *Parser {
	return &Parser{nil}
}

func (p *Parser) Parse(monkeyObservationDataLines string) *Monkey {
	lines := strings.Split(monkeyObservationDataLines, "\n")
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
		case testLineRegex.MatchString(line):
			testStrings := testLineRegex.FindStringSubmatch(line)
			divisor, _ := strconv.Atoi(testStrings[1])
			test := DivisibleBy(divisor)
			p.currentMonkeyBuilder.Test(test)
		case actionLineRegex.MatchString(line):
			actionStrings := actionLineRegex.FindStringSubmatch(line)
			condition := actionStrings[1]
			action := ThrowTo(actionStrings[2])
			if condition == "true" {
				p.currentMonkeyBuilder.WhenTrue(action)
			} else {
				p.currentMonkeyBuilder.WhenFalse(action)
			}
		}
	}
	return p.currentMonkeyBuilder.Build()
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
