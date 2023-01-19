package main

import (
	"log"
	"math/big"
)

type MonkeyRepository struct {
	monkeyIds    []string
	knownMonkeys map[string]*Monkey
	modulo       *big.Int
}

func (mr *MonkeyRepository) Add(m *Monkey) {
	mr.monkeyIds = append(mr.monkeyIds, m.id)
	mr.knownMonkeys[m.id] = m
	mr.modulo.Mul(mr.modulo, m.divisor)
}

func (mr *MonkeyRepository) Get(id string) *Monkey {
	return mr.knownMonkeys[id]
}

func (mr *MonkeyRepository) AllMonkeyIds() []string {
	return mr.monkeyIds
}

func newMonkeyRepository() *MonkeyRepository {
	return &MonkeyRepository{
		monkeyIds:    []string{},
		knownMonkeys: make(map[string]*Monkey),
		modulo:       big.NewInt(1),
	}
}
func GetMonkeyRepository() *MonkeyRepository {
	return newMonkeyRepository()
}

func (mr *MonkeyRepository) ThrowTo(monkeyId string) Action {
	return func(in *big.Int) {
		if in.Cmp(mr.modulo) == 1 {
			_, m := in.DivMod(in, mr.modulo, big.NewInt(0))
			in = m
		}
		monkey := mr.Get(monkeyId)
		if monkey != nil {
			monkey.AddItem(in)
			return
		}
		log.Printf("AllMonkeys : %#v", mr)
		log.Fatalf("Monkey with id %s not found", monkeyId)
	}
}
