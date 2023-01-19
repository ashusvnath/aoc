package main

import (
	"math/big"
	"runtime"
	"testing"
)

func assertEqual[T comparable](expected, actual T, t *testing.T) {
	_, file, line, _ := runtime.Caller(1)
	if expected != actual {
		t.Errorf("\n%s:%d\nExpected : %v\nActual   : %v\n", file, line, expected, actual)
	}
}

func assertEqualBigInt(expected, actual *big.Int, t *testing.T) {
	_, file, line, _ := runtime.Caller(1)
	if expected.Cmp(actual) != 0 {
		t.Errorf("\n%s:%d\nExpected : %v\nActual   : %v\n", file, line, expected, actual)
	}
}

func assertTrue(in bool, t *testing.T) {
	_, file, line, _ := runtime.Caller(1)
	if !in {
		t.Errorf("\n%s:%d\nExpected : %v to be true", file, line, in)
	}
}

func assertFalse(in bool, t *testing.T) {
	_, file, line, _ := runtime.Caller(1)
	if in {
		t.Errorf("%s:%d\nExpected : %v to be false", file, line, in)
	}
}
