package assert

import (
	"runtime"
	"testing"
)

//lint:ignore U1000 this is a utility
func Equal[T comparable](expected, actual T, t *testing.T) {
	_, file, line, _ := runtime.Caller(1)
	if expected != actual {
		t.Errorf("\n%s:%d\nExpected : %v\nActual   : %v\n", file, line, expected, actual)
	}
}

//lint:ignore U1000 this is a utility
func True(in bool, t *testing.T) {
	_, file, line, _ := runtime.Caller(1)
	if !in {
		t.Errorf("\n%s:%d\nExpected : %v to be true", file, line, in)
	}
}

//lint:ignore U1000 this is a utility
func False(in bool, t *testing.T) {
	_, file, line, _ := runtime.Caller(1)
	if in {
		t.Errorf("%s:%d\nExpected : %v to be false", file, line, in)
	}
}
