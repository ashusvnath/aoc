package main

import "testing"

func TestParseOperation(t *testing.T) {
	t.Run("should parse multiplication", func(t *testing.T) {
		input := []string{"  Operation: new = old * 19", "new = old * 19", "*", "19"}
		operation := ParseOperation(input)
		assertEqual(38, operation(2), t)
	})

	t.Run("should parse addition", func(t *testing.T) {
		input := []string{"  Operation: new = old + 6", "new = old + 6", "+", "6"}
		operation := ParseOperation(input)
		assertEqual(8, operation(2), t)
	})

	t.Run("should parse square", func(t *testing.T) {
		input := []string{"  Operation: new = old * old", "new = old * old", "*", "old"}
		operation := ParseOperation(input)
		assertEqual(81, operation(9), t)
	})

	t.Run("Should parse square end to end", func(t *testing.T) {
		operation := ParseOperation(operationRegex.FindStringSubmatch("  Operation: new = old * old"))
		assertEqual(81, operation(9), t)
	})
}

func TestParse(t *testing.T) {
	monkey2 := NewMonkeyBuilder().Id("2").Build()
	monkey3 := NewMonkeyBuilder().Id("3").Build()
	input := `Monkey 0:
	Starting items: 79, 69
	Operation: new = old * 19
	Test: divisible by 23
	  If true: throw to monkey 2
	  If false: throw to monkey 3`
	monkey0 := NewParser().Parse(input)

	repo.Add(monkey0)
	repo.Add(monkey2)
	repo.Add(monkey3)

	assertEqual(79, monkey0.items[0], t)
	assertEqual(69, monkey0.items[1], t)
	assertEqual(38, monkey0.op(2), t)

	monkey0.DoBusiness()
	assertEqual(500, monkey3.items[0], t)
	assertEqual(437, monkey2.items[0], t)
}
