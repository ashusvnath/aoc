package main

import "log"

func init() {
	if __allMonkeys == nil {
		__allMonkeys = make(map[string]*Monkey)
	}
}

var __allMonkeys map[string]*Monkey

func __getAllMonkeys() map[string]*Monkey {
	return __allMonkeys
}

type Monkey struct {
	id        string
	op        Operation
	action    Action
	items     []int
	processed []int
}

type MonkeyBuilder struct {
	monkey *Monkey
}

func (mb *MonkeyBuilder) Id(id string) *MonkeyBuilder {
	mb.monkey.id = id
	return mb
}

func (mb *MonkeyBuilder) Op(op Operation) *MonkeyBuilder {
	mb.monkey.op = op
	return mb
}

func (mb *MonkeyBuilder) Action(action Action) *MonkeyBuilder {
	mb.monkey.action = action
	return mb
}

func (mb *MonkeyBuilder) Items(items ...int) *MonkeyBuilder {
	mb.monkey.items = items
	return mb
}

func (m MonkeyBuilder) Build() *Monkey {
	__allMonkeys[m.monkey.id] = m.monkey
	return m.monkey
}

func (m *Monkey) DoBusiness() {
	for _, item := range m.items {
		m.action(Divide(3)(m.op(item)))
		m.processed = append(m.processed, item)
	}
	m.items = nil
}

func (m *Monkey) AddItem(n int) {
	log.Printf("Monkey%s: adding item %v", m.id, n)
	m.items = append(m.items, n)
}

func (m *Monkey) Activity() int {
	return len(m.processed)
}

func ThrowTo(monkeyId string) Action {
	return func(in int) {
		monkey := __getAllMonkeys()[monkeyId]
		if monkey != nil {
			monkey.AddItem(in)
			return
		}
		log.Printf("AllMonkeys : %#v", __getAllMonkeys())
		log.Fatalf("Monkey with id %s not found", monkeyId)
	}
}

func NewMonkeyBuilder() *MonkeyBuilder {
	return &MonkeyBuilder{&Monkey{"unassigned", nil, nil, nil, nil}}
}
