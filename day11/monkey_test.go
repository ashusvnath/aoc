package main

import "testing"

func TestMonkey(t *testing.T) {

	t.Run("Monkeys business should be performed in order of items provided", func(t *testing.T) {
		m1 := NewMonkeyBuilder().Id("1").Build()
		m2 := NewMonkeyBuilder().Id("2").Build()

		m3 := NewMonkeyBuilder().
			Id("3").
			Op(Mul(10)).
			Action(Conditional(DivisibleBy(7), ThrowTo("1"), ThrowTo("2"))).
			Items(21, 36).
			Build()

		m3.DoBusiness()
		assertEqual(2, m3.Activity(), t)
		assertEqual(0, len(m3.items), t)
		assertEqual(70, m1.items[0], t)
		assertEqual(120, m2.items[0], t)
	})

}
