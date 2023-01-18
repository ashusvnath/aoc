package main

import "testing"

func TestOperations(t *testing.T) {
	t.Run("Add(6) should add 6", func(t *testing.T) {
		op := Add(6)
		assertEqual(10, op(4), t)
	})

	t.Run("Mul(7) should mulipty by 7", func(t *testing.T) {
		op := Mul(7)
		assertEqual(28, op(4), t)
	})

	t.Run("Square should multiply itself", func(t *testing.T) {
		op := Square()
		assertEqual(16, op(4), t)
	})

	t.Run("Divide(3) should perform integer division by 3", func(t *testing.T) {
		op := Divide(3)
		assertEqual(10, op(31), t)
		assertEqual(7, op(21), t)
		assertEqual(7, op(22), t)
		assertEqual(7, op(23), t)
	})
}

func TestTests(t *testing.T) {
	t.Run("DivisibleBy(5) should check if input is dividsible by 5", func(t *testing.T) {
		test := DivisibleBy(5)
		assertTrue(test(10), t)
		assertFalse(test(11), t)
	})
}

func TestConditional(t *testing.T) {
	test := DivisibleBy(23)
	failTestAction := func(_ int) { t.Error("this action should not have been called") }
	createPassTestAction := func(x int) Action {
		return func(i int) {
			assertEqual(x, i, t)
		}
	}

	t.Run("Should execute true action when test passes", func(t *testing.T) {
		input := 46
		cond := Conditional(test, createPassTestAction(input), failTestAction)
		cond(input)
	})

	t.Run("Should execute false action when test fails", func(t *testing.T) {
		input := 47
		cond := Conditional(test, failTestAction, createPassTestAction(input))
		cond(input)
	})
}
