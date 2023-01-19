package main

import (
	"math/big"
	"testing"
)

func TestOperations(t *testing.T) {
	t.Run("Add(6) should add 6", func(t *testing.T) {
		op := Add(big.NewInt(6))
		assertEqualBigInt(big.NewInt(10), op(big.NewInt(4)), t)
	})

	t.Run("Mul(7) should mulipty by 7", func(t *testing.T) {
		op := Mul(big.NewInt(7))
		assertEqualBigInt(big.NewInt(28), op(big.NewInt(4)), t)
	})

	t.Run("Square should multiply itself", func(t *testing.T) {
		op := Square()
		assertEqualBigInt(big.NewInt(16), op(big.NewInt(4)), t)
	})

	t.Run("Divide(3) should perform integer division by 3", func(t *testing.T) {
		op := Divide(big.NewInt(3))
		assertEqualBigInt(big.NewInt(10), op(big.NewInt(31)), t)
		assertEqualBigInt(big.NewInt(7), op(big.NewInt(21)), t)
		assertEqualBigInt(big.NewInt(7), op(big.NewInt(22)), t)
		assertEqualBigInt(big.NewInt(7), op(big.NewInt(23)), t)
	})
}

func TestTests(t *testing.T) {
	t.Run("DivisibleBy(5) should check if input is dividsible by 5", func(t *testing.T) {
		test := DivisibleBy(big.NewInt(5))
		assertTrue(test(big.NewInt(10)), t)
		assertFalse(test(big.NewInt(11)), t)
	})
}

func TestConditional(t *testing.T) {
	test := DivisibleBy(big.NewInt(23))
	failTestAction := func(_ *big.Int) { t.Error("this action should not have been called") }
	createPassTestAction := func(x *big.Int) Action {
		return func(i *big.Int) {
			assertEqualBigInt(x, i, t)
		}
	}

	t.Run("Should execute true action when test passes", func(t *testing.T) {
		input := big.NewInt(46)
		cond := Conditional(test, createPassTestAction(input), failTestAction)
		cond(input)
	})

	t.Run("Should execute false action when test fails", func(t *testing.T) {
		input := big.NewInt(47)
		cond := Conditional(test, failTestAction, createPassTestAction(input))
		cond(input)
	})
}
