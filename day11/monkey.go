package main

import (
	"fmt"
	"math/big"
)

type Monkey struct {
	id        string
	op        Operation
	action    Action
	items     []*big.Int
	processed int
	divisor   *big.Int
}

func (m *Monkey) DoBusiness() {
	for _, item := range m.items {
		m.action(m.op(item))
		m.processed++
	}
	m.items = nil
}

func (m *Monkey) AddItem(n *big.Int) {
	m.items = append(m.items, n)
}

func (m *Monkey) Activity() int {
	return m.processed
}

func (m *Monkey) GoString() string {
	return fmt.Sprintf("Monkey%s(activity: %v)", m.id, m.Activity())
}

func (m *Monkey) String() string {
	return fmt.Sprintf("Monkey%s(activity: %v)", m.id, m.Activity())
}
