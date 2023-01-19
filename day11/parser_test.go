package main

import (
	"math/big"
	"testing"
)

func TestParseOperation(t *testing.T) {
	t.Run("should parse multiplication", func(t *testing.T) {
		input := []string{"  Operation: new = old * 19", "new = old * 19", "*", "19"}
		operation := ParseOperation(input)
		assertEqualBigInt(big.NewInt(38), operation(big.NewInt(2)), t)
	})

	t.Run("should parse addition", func(t *testing.T) {
		input := []string{"  Operation: new = old + 6", "new = old + 6", "+", "6"}
		operation := ParseOperation(input)
		assertEqualBigInt(big.NewInt(8), operation(big.NewInt(2)), t)
	})

	t.Run("should parse square", func(t *testing.T) {
		input := []string{"  Operation: new = old * old", "new = old * old", "*", "old"}
		operation := ParseOperation(input)
		assertEqualBigInt(big.NewInt(81), operation(big.NewInt(9)), t)
	})

	t.Run("Should parse square end to end", func(t *testing.T) {
		operation := ParseOperation(operationRegex.FindStringSubmatch("  Operation: new = old * old"))
		assertEqualBigInt(big.NewInt(81), operation(big.NewInt(9)), t)
	})
}

func TestParse(t *testing.T) {
	repo := GetMonkeyRepository()
	monkey2 := NewMonkeyBuilder(Identity()).Id("2").Divisor(big.NewInt(100)).Build()
	monkey3 := NewMonkeyBuilder(Identity()).Id("3").Divisor(big.NewInt(100)).Build()
	input := `Monkey 0:
	Starting items: 79, 69
	Operation: new = old * 19
	Test: divisible by 23
	  If true: throw to monkey 2
	  If false: throw to monkey 3`
	monkey0 := NewParser(Divide(big.NewInt(3)), repo).Parse(input)

	repo.Add(monkey0)
	repo.Add(monkey2)
	repo.Add(monkey3)

	assertEqualBigInt(big.NewInt(79), monkey0.items[0], t)
	assertEqualBigInt(big.NewInt(69), monkey0.items[1], t)
	assertEqualBigInt(big.NewInt(12), monkey0.op(big.NewInt(2)), t)

	monkey0.DoBusiness()
	assertEqualBigInt(big.NewInt(500), monkey3.items[0], t)
	assertEqualBigInt(big.NewInt(437), monkey2.items[0], t)
}
