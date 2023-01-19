package main

import "log"

type MonkeyRepository struct {
	knownMonkeys map[string]*Monkey
}

var repo *MonkeyRepository = newMonkeyRepository()

func (mr *MonkeyRepository) Add(m *Monkey) {
	mr.knownMonkeys[m.id] = m
}

func (mr *MonkeyRepository) Get(id string) *Monkey {
	return mr.knownMonkeys[id]
}

func newMonkeyRepository() *MonkeyRepository {
	return &MonkeyRepository{make(map[string]*Monkey)}
}
func GetMonkeyRepository() *MonkeyRepository {
	return repo
}

func ThrowTo(monkeyId string) Action {
	return func(in int) {
		monkey := repo.Get(monkeyId)
		if monkey != nil {
			monkey.AddItem(in)
			return
		}
		log.Printf("AllMonkeys : %#v", repo)
		log.Fatalf("Monkey with id %s not found", monkeyId)
	}
}
