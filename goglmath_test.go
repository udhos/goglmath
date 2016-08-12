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
	m1 := &Matrix4{[16]float32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}}
	if !m1.Identity() {
		t.Errorf("not identity matrix")
	}
}

func TestIdentity2(t *testing.T) {
	m1 := &Matrix4{[16]float32{
		1, 0, 0, 1,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}}
	if m1.Identity() {
		t.Errorf("wrong identity matrix")
	}
}

func TestIdentity3(t *testing.T) {
	m1 := &Matrix4{[16]float32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 1,
	}}
	if m1.Identity() {
		t.Errorf("wrong identity matrix")
	}
}

func TestIdentity4(t *testing.T) {
	m1 := &Matrix4{}
	m1.SetIdentity()
	if !m1.Identity() {
		t.Errorf("not identity matrix")
	}
}

func TestIdentity5(t *testing.T) {
	m1 := NewMatrix4Identity()
	if !m1.Identity() {
		t.Errorf("not identity matrix")
	}
}

func BenchmarkSetNullMatrix(b *testing.B) {
	m := &Matrix4{}
	for n := 0; n < b.N; n++ {
		m.SetNull()
	}
}

func BenchmarkSetIdentityMatrix(b *testing.B) {
	m := &Matrix4{}
	for n := 0; n < b.N; n++ {
		m.SetIdentity()
	}
}

func BenchmarkLengthSquared3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lengthSquared3(1, 2, 3)
	}
}

func BenchmarkIdentity(b *testing.B) {
	m1 := Matrix4{}
	for n := 0; n < b.N; n++ {
		m1.Identity()
	}
}

func BenchmarkNull(b *testing.B) {
	m1 := Matrix4{}
	for n := 0; n < b.N; n++ {
		m1.Null()
	}
}
