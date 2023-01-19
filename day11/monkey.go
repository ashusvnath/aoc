package main

import (
	"fmt"
	"log"
)

type Monkey struct {
	id        string
	op        Operation
	action    Action
	items     []int
	processed []int
}

func (m *Monkey) DoBusiness() {
	log.Printf("Monkey%s: Doing business", m.id)
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

func (m *Monkey) GoString() string {
	return fmt.Sprintf("Monkey%s(activity: %v)", m.id, m.Activity()) //, m.processed
}

func (m *Monkey) String() string {
	return fmt.Sprintf("Monkey%s(activity: %v)", m.id, m.Activity()) //, m.processed
}

func NewMonkeyBuilder() *MonkeyBuilder {
	return &MonkeyBuilder{
		monkey: &Monkey{"unassigned", nil, nil, nil, nil},
		cb:     &ConditionalBuilder{nil, nil, nil},
	}
}
