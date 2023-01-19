package main

import (
	"log"
	"math/big"
	"regexp"
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
	monkeyBuilder *MonkeyBuilder
	decoration    Operation
	repo          *MonkeyRepository
}

func NewParser(decoration Operation, repo *MonkeyRepository) *Parser {
	return &Parser{nil, decoration, repo}
}

func (p *Parser) Parse(monkeyObservationDataLines string) *Monkey {
	lines := strings.Split(monkeyObservationDataLines, "\n")
	p.monkeyBuilder = NewMonkeyBuilder(p.decoration)
	for _, line := range lines {
		switch {
		case monkeyLineRegex.MatchString(line):
			id := monkeyLineRegex.FindStringSubmatch(line)[1]
			p.monkeyBuilder.Id(id)
		case itemsLineRegex.MatchString(line):
			itemsString := itemsLineRegex.FindStringSubmatch(line)[1]
			itemsAsStrings := strings.Split(itemsString, ", ")
			items := make([]*big.Int, len(itemsAsStrings))
			for i, item := range itemsAsStrings {
				items[i], _ = big.NewInt(0).SetString(item, 10)
			}
			p.monkeyBuilder.Items(items...)
		case operationRegex.MatchString(line):
			operationStrings := operationRegex.FindStringSubmatch(line)
			operation := ParseOperation(operationStrings)
			if p.decoration != nil {
				t := operation
				operation = func(in *big.Int) *big.Int { return p.decoration(t(in)) }
			}
			p.monkeyBuilder.Op(operation)
		case testLineRegex.MatchString(line):
			testStrings := testLineRegex.FindStringSubmatch(line)
			divisor, _ := big.NewInt(0).SetString(testStrings[1], 0)
			test := DivisibleBy(divisor)
			p.monkeyBuilder.Test(test)
			p.monkeyBuilder.Divisor(divisor)
		case actionLineRegex.MatchString(line):
			actionStrings := actionLineRegex.FindStringSubmatch(line)
			condition := actionStrings[1]
			action := p.repo.ThrowTo(actionStrings[2])
			if condition == "true" {
				p.monkeyBuilder.WhenTrue(action)
			} else {
				p.monkeyBuilder.WhenFalse(action)
			}
		}
	}
	return p.monkeyBuilder.Build()
}

func ParseOperation(input []string) Operation {
	operation := input[2]
	operand := input[3]
	switch operation {
	case "*":
		if operand == "old" {
			return Square()
		}
		opInt, _ := big.NewInt(0).SetString(operand, 0)
		return Mul(opInt)
	case "+":
		opInt, _ := big.NewInt(0).SetString(operand, 0)
		return Add(opInt)
	default:
		log.Fatalf("Unknown operation %s", operation)
	}
	return nil
}
