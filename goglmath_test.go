package goglmath

import (
	"testing"
)

func TestNormalize3(t *testing.T) {
	a, b, c := Normalize3(1, 2, 3)
	size := Length3(a, b, c)
	want := 1.0
	if size != want {
		t.Errorf("expected=%v got=%v", want, size)
	}
}
