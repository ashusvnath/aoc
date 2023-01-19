package main

import (
	"math/big"
	"testing"
)

func TestMonkey(t *testing.T) {

	t.Run("Monkeys business should be performed in order of provided items", func(t *testing.T) {
		repo := GetMonkeyRepository()
		m1 := NewMonkeyBuilder(Identity()).Id("1").Divisor(big.NewInt(100)).Build()
		m2 := NewMonkeyBuilder(Identity()).Id("2").Divisor(big.NewInt(100)).Build()

		m3 := NewMonkeyBuilder(Divide(big.NewInt(3))).
			Id("3").
			Op(Mul(big.NewInt(10))).
			Test(DivisibleBy(big.NewInt(7))).
			Divisor(big.NewInt(7)).
			WhenTrue(repo.ThrowTo("1")).
			WhenFalse(repo.ThrowTo("2")).
			Items(big.NewInt(21), big.NewInt(36)).
			Build()

		repo.Add(m1)
		repo.Add(m2)
		repo.Add(m3)
		//repo.modulo = big.NewInt(100000)

		m3.DoBusiness()
		assertEqual(2, m3.Activity(), t)
		assertEqual(0, len(m3.items), t)
		assertEqualBigInt(big.NewInt(210), m1.items[0], t)
		assertEqualBigInt(big.NewInt(360), m2.items[0], t)
	})
}
