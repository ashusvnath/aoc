package main

import (
	"testing"
)

func assertEqual[T comparable](expected, actual T, t *testing.T) {
	if expected != actual {
		t.Errorf("\nExpected : %v\nActual   : %v\n", expected, actual)
	}
}

func assertTrue(in bool, t *testing.T) {
	if !in {
		t.Errorf("\nExpected : %v to be true", in)
	}
}

func assertFalse(in bool, t *testing.T) {
	if in {
		t.Errorf("\nExpected : %v to be false", in)
	}
}
