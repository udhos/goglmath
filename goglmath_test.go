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
	m2.CopyFrom(m1)
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

func TestNull1(t *testing.T) {
	m1 := &Matrix4{}
	if !m1.Null() {
		t.Errorf("null matrix reported as non-null: %v", m1)
	}
}

func TestNull2(t *testing.T) {
	m1 := &Matrix4{[16]float32{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 1,
	}}
	if m1.Null() {
		t.Errorf("non-null matrix reported as null")
	}
}

func TestNull3(t *testing.T) {
	m1 := &Matrix4{}
	m1.SetIdentity()
	if !m1.Identity() {
		t.Errorf("identity matrix reported as non-identity")
	}
	if m1.Null() {
		t.Errorf("identity matrix reported as null")
	}
	m1.SetNull()
	if m1.Identity() {
		t.Errorf("null matrix reported as identity")
	}
	if !m1.Null() {
		t.Errorf("null matrix reported as non-null: %v", m1)
	}
}

func BenchmarkData(b *testing.B) {
	m := &Matrix4{}
	for n := 0; n < b.N; n++ {
		d := m.Data()
		d[0]++
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
		LengthSquared3(1, 2, 3)
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
