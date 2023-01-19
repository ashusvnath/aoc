package main

import (
	"sort"
	"strings"
)

type MonkeyService struct {
	repo   *MonkeyRepository
	parser *Parser
}

func (ms *MonkeyService) Setup(input string) {
	monkeyObservationData := strings.Split(input, "\n\n")
	for _, monkeyObservationLines := range monkeyObservationData {
		ms.repo.Add(ms.parser.Parse(monkeyObservationLines))
	}
}

func (ms *MonkeyService) DoBusiness(rounds int) {
	ids := ms.repo.AllMonkeyIds()
	for i := 0; i < rounds; i++ {
		for _, id := range ids {
			repo.Get(id).DoBusiness()
		}

	}
}

func (ms *MonkeyService) MonkeysByActivity() []*Monkey {
	monkeys := []*Monkey{}
	for _, id := range ms.repo.AllMonkeyIds() {
		monkeys = append(monkeys, ms.repo.Get(id))
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Activity() > monkeys[j].Activity()
	})
	return monkeys
}

func NewMonkeyService() *MonkeyService {
	return &MonkeyService{
		repo:   GetMonkeyRepository(),
		parser: NewParser(),
	}
}
