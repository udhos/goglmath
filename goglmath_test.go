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

func TestCopy(t *testing.T) {
	m1 := &Matrix4{}
	m2 := &Matrix4{}
	for i := 0; i < 16; i++ {
		m1.data[i] = float32(i)
	}
	m2.copyFrom(m1)
	for i := 0; i < 16; i++ {
		if m1.data[i] != m2.data[i] {
			t.Errorf("matrix copy fail: position %d: expected=%v got=%v", i, m1.data[i], m2.data[i])
		}
	}
}

func TestIdentity1(t *testing.T) {
	m1 := &Matrix4{}
	SetIdentityMatrix(m1)
	if !isIdentityMatrix(m1) {
		t.Errorf("not identity matrix")
	}
}

func TestIdentity2(t *testing.T) {
	m1 := NewMatrix4Identity()
	if !isIdentityMatrix(m1) {
		t.Errorf("not identity matrix")
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
