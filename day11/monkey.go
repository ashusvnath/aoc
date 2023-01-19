package main

import "log"

type Monkey struct {
	id        string
	op        Operation
	action    Action
	items     []int
	processed []int
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

func NewMonkeyBuilder() *MonkeyBuilder {
	return &MonkeyBuilder{
		monkey: &Monkey{"unassigned", nil, nil, nil, nil},
		cb:     &ConditionalBuilder{nil, nil, nil},
	}
}
