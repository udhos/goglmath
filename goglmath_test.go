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

func BenchmarkSetNullMatrix(b *testing.B) {
	m := &Matrix4{}
	//m.Malloc()

	for n := 0; n < b.N; n++ {
		SetNullMatrix(m)
	}
}

func BenchmarkLengthSquared3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lengthSquared3(1, 2, 3)
	}
}
