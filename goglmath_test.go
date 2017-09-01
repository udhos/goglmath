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

func TestOrthoMatrix(t *testing.T) {
	var O Matrix4
	SetOrthoMatrix(&O, -1, 1, -1, 1, 1, -1)
	I := NewMatrix4Identity()
	if !Matrix4Equal(&O, &I) {
		t.Errorf("mismatch: ortho=%v identity=%v", O, I)
	}
}

// forward = 0 0 -1 // looking towards -Z
// up = 0 1 0       // up direction is +Y
func TestRotate(t *testing.T) {
	M := NewMatrix4Identity()
	M.Rotate(0, 0, -1, 0, -1, 0) // turn upside down (180 around Z)
	px, py, pz, pw := M.Transform(0, 1, 0, 1)
	if px != 0 || py != -1 || pz != 0 || pw != 1 {
		t.Errorf("unexpected point location: %f,%f,%f,%f", px, py, pz, pw)
	}
}

func BenchmarkMatrix4Equal1(b *testing.B) {
	m := NewMatrix4Identity()
	for n := 0; n < b.N; n++ {
		matrix4Equal1(&m, &m)
	}
}

func BenchmarkMatrix4Equal2(b *testing.B) {
	m := NewMatrix4Identity()
	for n := 0; n < b.N; n++ {
		matrix4Equal2(&m, &m)
	}
}

func BenchmarkData(b *testing.B) {
	m := &Matrix4{}
	for n := 0; n < b.N; n++ {
		d := m.Data()
		d[0]++
	}
}

func BenchmarkSetNullMatrixItems(b *testing.B) {
	m := &Matrix4{}
	for n := 0; n < b.N; n++ {
		m.setNullItems()
	}
}

func BenchmarkSetNullMatrixCopy(b *testing.B) {
	m := &Matrix4{}
	for n := 0; n < b.N; n++ {
		m.setNullCopy()
	}
}

func BenchmarkSetNullMatrix(b *testing.B) {
	m := &Matrix4{}
	for n := 0; n < b.N; n++ {
		m.SetNull()
	}
}

func BenchmarkSetIdentityMatrixItems(b *testing.B) {
	m := &Matrix4{}
	for n := 0; n < b.N; n++ {
		m.setIdentityItems()
	}
}

func BenchmarkSetIdentityMatrixCopy(b *testing.B) {
	m := &Matrix4{}
	for n := 0; n < b.N; n++ {
		m.setIdentityCopy()
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

func BenchmarkNullRange(b *testing.B) {
	m1 := Matrix4{}
	for n := 0; n < b.N; n++ {
		m1.nullRange()
	}
}

func BenchmarkNullComp(b *testing.B) {
	m1 := Matrix4{}
	for n := 0; n < b.N; n++ {
		m1.nullComp()
	}
}

func BenchmarkNull(b *testing.B) {
	m1 := Matrix4{}
	for n := 0; n < b.N; n++ {
		m1.Null()
	}
}

func BenchmarkViewportTransform1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		viewportTransform1(10, 10, 100, 100, .1, .9, 1.0, 2.0, 3.0)
	}
}

func BenchmarkViewportTransform2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		viewportTransform2(10, 10, 100, 100, .1, .9, 1.0, 2.0, 3.0)
	}
}

func BenchmarkViewportTransform(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ViewportTransform(10, 10, 100, 100, .1, .9, 1.0, 2.0, 3.0)
	}
}
