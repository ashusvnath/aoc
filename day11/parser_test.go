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
