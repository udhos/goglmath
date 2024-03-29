// Package goglmath is a lightweight pure Go 3D math package providing essential matrix/vector operations for GL graphics applications.
package goglmath

import (
	"errors"
	"math"
	"reflect"
)

// Matrix4 is a 4x4 matrix.
type Matrix4 struct {
	data [16]float32
}

var mat4identity = Matrix4{[16]float32{
	1, 0, 0, 0,
	0, 1, 0, 0,
	0, 0, 1, 0,
	0, 0, 0, 1,
}}

var mat4null = Matrix4{}

// NewMatrix4Identity creates an identity matrix.
func NewMatrix4Identity() Matrix4 {
	return mat4identity // clone -- it's unsafe to return pointer to the original data
}

// NewMatrix4Null creates a null matrix.
func NewMatrix4Null() Matrix4 {
	return mat4null // clone -- it's unsafe to return pointer to the original data
}

// Data returns matrix data as slice ready to GPU upload.
func (m *Matrix4) Data() []float32 {
	return m.data[:]
}

// Matrix4Equal tests matrices for equality.
func Matrix4Equal(m1, m2 *Matrix4) bool {
	return matrix4Equal3(m1, m2)
}

func matrix4Equal1(m1, m2 *Matrix4) bool {
	for i, v := range m1.data {
		if v != m2.data[i] {
			return false
		}
	}
	return true
}

func matrix4Equal2(m1, m2 *Matrix4) bool {
	return reflect.DeepEqual(m1.data, m2.data)
}

func matrix4Equal3(m1, m2 *Matrix4) bool {
	// Array values are comparable if values of the array element type are comparable. Two array values are equal if their corresponding elements are equal.
	return m1.data == m2.data
}

// Identity reports if matrix is identity.
func (m *Matrix4) Identity() bool {
	// Array values are comparable if values of the array element type are comparable. Two array values are equal if their corresponding elements are equal.
	return m.data == mat4identity.data
	//return Matrix4Equal(m, &mat4identity)
}

// Null reports if matrix is null.
func (m *Matrix4) Null() bool {
	return m.nullComp()
}

func (m *Matrix4) nullRange() bool {
	for _, f := range m.data {
		if f != 0 {
			return false
		}
	}
	return true
}

func (m *Matrix4) nullComp() bool {
	// Array values are comparable if values of the array element type are comparable. Two array values are equal if their corresponding elements are equal.
	return m.data == mat4null.data
	//return Matrix4Equal(m, &mat4null)
}

// CopyFrom copy matrix data from another source matrix.
func (m *Matrix4) CopyFrom(src *Matrix4) {
	// Go's arrays are values. An array variable denotes the entire array; it is not a pointer to the first array element (as would be the case in C). This means that when you assign or pass around an array value you will make a copy of its contents.
	m.data = src.data
}

// SetNull sets matrix to null.
func (m *Matrix4) SetNull() {
	m.setNullCopy()
}

func (m *Matrix4) setNullItems() {
	m.data[0] = 0
	m.data[1] = 0
	m.data[2] = 0
	m.data[3] = 0
	m.data[4] = 0
	m.data[5] = 0
	m.data[6] = 0
	m.data[7] = 0
	m.data[8] = 0
	m.data[9] = 0
	m.data[10] = 0
	m.data[11] = 0
	m.data[12] = 0
	m.data[13] = 0
	m.data[14] = 0
	m.data[15] = 0
}

func (m *Matrix4) setNullCopy() {
	// Go's arrays are values. An array variable denotes the entire array; it is not a pointer to the first array element (as would be the case in C). This means that when you assign or pass around an array value you will make a copy of its contents.
	m.data = mat4null.data
}

func (m *Matrix4) setIdentityItems() {
	m.data[0] = 1
	m.data[1] = 0
	m.data[2] = 0
	m.data[3] = 0
	m.data[4] = 0
	m.data[5] = 1
	m.data[6] = 0
	m.data[7] = 0
	m.data[8] = 0
	m.data[9] = 0
	m.data[10] = 1
	m.data[11] = 0
	m.data[12] = 0
	m.data[13] = 0
	m.data[14] = 0
	m.data[15] = 1
}

func (m *Matrix4) setIdentityCopy() {
	// Go's arrays are values. An array variable denotes the entire array; it is not a pointer to the first array element (as would be the case in C). This means that when you assign or pass around an array value you will make a copy of its contents.
	m.data = mat4identity.data
}

// SetIdentity sets matrix to identity.
func (m *Matrix4) SetIdentity() {
	m.setIdentityCopy()
}

// Invert inverts the matrix.
func (m *Matrix4) Invert() error {
	return m.CopyInverseFrom(m)
}

// CopyInverseFrom sets the matrix as inverse of another source matrix.
func (m *Matrix4) CopyInverseFrom(src *Matrix4) error {
	a00 := src.data[0]
	a01 := src.data[1]
	a02 := src.data[2]
	a03 := src.data[3]
	a10 := src.data[4]
	a11 := src.data[5]
	a12 := src.data[6]
	a13 := src.data[7]
	a20 := src.data[8]
	a21 := src.data[9]
	a22 := src.data[10]
	a23 := src.data[11]
	a30 := src.data[12]
	a31 := src.data[13]
	a32 := src.data[14]
	a33 := src.data[15]

	b00 := a00*a11 - a01*a10
	b01 := a00*a12 - a02*a10
	b02 := a00*a13 - a03*a10
	b03 := a01*a12 - a02*a11
	b04 := a01*a13 - a03*a11
	b05 := a02*a13 - a03*a12
	b06 := a20*a31 - a21*a30
	b07 := a20*a32 - a22*a30
	b08 := a20*a33 - a23*a30
	b09 := a21*a32 - a22*a31
	b10 := a21*a33 - a23*a31
	b11 := a22*a33 - a23*a32

	det := b00*b11 - b01*b10 + b02*b09 + b03*b08 - b04*b07 + b05*b06
	if det == 0.0 {
		m.CopyFrom(src)
		return errors.New("copyInverseFrom: null determinant")
	}
	invDet := 1.0 / det

	m.data[0] = (a11*b11 - a12*b10 + a13*b09) * invDet
	m.data[1] = (-a01*b11 + a02*b10 - a03*b09) * invDet
	m.data[2] = (a31*b05 - a32*b04 + a33*b03) * invDet
	m.data[3] = (-a21*b05 + a22*b04 - a23*b03) * invDet
	m.data[4] = (-a10*b11 + a12*b08 - a13*b07) * invDet
	m.data[5] = (a00*b11 - a02*b08 + a03*b07) * invDet
	m.data[6] = (-a30*b05 + a32*b02 - a33*b01) * invDet
	m.data[7] = (a20*b05 - a22*b02 + a23*b01) * invDet
	m.data[8] = (a10*b10 - a11*b08 + a13*b06) * invDet
	m.data[9] = (-a00*b10 + a01*b08 - a03*b06) * invDet
	m.data[10] = (a30*b04 - a31*b02 + a33*b00) * invDet
	m.data[11] = (-a20*b04 + a21*b02 - a23*b00) * invDet
	m.data[12] = (-a10*b09 + a11*b07 - a12*b06) * invDet
	m.data[13] = (a00*b09 - a01*b07 + a02*b06) * invDet
	m.data[14] = (-a30*b03 + a31*b01 - a32*b00) * invDet
	m.data[15] = (a20*b03 - a21*b01 + a22*b00) * invDet

	return nil
}

// Transform multiples this matrix [m] by vector [x,y,z,w]
func (m *Matrix4) Transform(x, y, z, w float64) (tx, ty, tz, tw float64) {
	m0 := float64(m.data[0])
	m1 := float64(m.data[1])
	m2 := float64(m.data[2])
	m3 := float64(m.data[3])
	m4 := float64(m.data[4])
	m5 := float64(m.data[5])
	m6 := float64(m.data[6])
	m7 := float64(m.data[7])
	m8 := float64(m.data[8])
	m9 := float64(m.data[9])
	m10 := float64(m.data[10])
	m11 := float64(m.data[11])
	m12 := float64(m.data[12])
	m13 := float64(m.data[13])
	m14 := float64(m.data[14])
	m15 := float64(m.data[15])

	tx = m0*x + m4*y + m8*z + m12*w
	ty = m1*x + m5*y + m9*z + m13*w
	tz = m2*x + m6*y + m10*z + m14*w
	tw = m3*x + m7*y + m11*z + m15*w

	return
}

// Rotate multiplies the matrix m by a rotation matrix built from specified forward and up vectors.
// This rotation will rotate a point from the "null rotation" direction to the direction specified by forward and up vectors.
// null rotation:
// forward = 0 0 -1 // looking towards -Z
// up = 0 1 0       // up direction is +Y
func (m *Matrix4) Rotate(forwardX, forwardY, forwardZ, upX, upY, upZ float64) {
	var rotate Matrix4
	SetRotationMatrix(&rotate, forwardX, forwardY, forwardZ, upX, upY, upZ)
	m.Multiply(&rotate)
}

// Translate multiples the matrix by a translation matrix.
// usually set w to 1.0
func (m *Matrix4) Translate(tx, ty, tz, tw float64) {
	x := float32(tx)
	y := float32(ty)
	z := float32(tz)
	w := float32(tw)
	t1 := m.data[0]*x + m.data[4]*y + m.data[8]*z + m.data[12]*w
	t2 := m.data[1]*x + m.data[5]*y + m.data[9]*z + m.data[13]*w
	t3 := m.data[2]*x + m.data[6]*y + m.data[10]*z + m.data[14]*w
	t4 := m.data[3]*x + m.data[7]*y + m.data[11]*z + m.data[15]*w
	m.data[12] = t1
	m.data[13] = t2
	m.data[14] = t3
	m.data[15] = t4
}

// Scale multiplies the matrix by a scaling matrix.
// usually set w to 1.0
func (m *Matrix4) Scale(x, y, z, w float64) {
	x1 := float32(x)
	y1 := float32(y)
	z1 := float32(z)
	w1 := float32(w)

	m.data[0] *= x1
	m.data[1] *= x1
	m.data[2] *= x1
	m.data[3] *= x1

	m.data[4] *= y1
	m.data[5] *= y1
	m.data[6] *= y1
	m.data[7] *= y1

	m.data[8] *= z1
	m.data[9] *= z1
	m.data[10] *= z1
	m.data[11] *= z1

	m.data[12] *= w1
	m.data[13] *= w1
	m.data[14] *= w1
	m.data[15] *= w1
}

// Multiply multiplies the matrix by another matrix.
func (m *Matrix4) Multiply(n *Matrix4) {
	m00 := m.data[0]
	m01 := m.data[4]
	m02 := m.data[8]
	m03 := m.data[12]
	m10 := m.data[1]
	m11 := m.data[5]
	m12 := m.data[9]
	m13 := m.data[13]
	m20 := m.data[2]
	m21 := m.data[6]
	m22 := m.data[10]
	m23 := m.data[14]
	m30 := m.data[3]
	m31 := m.data[7]
	m32 := m.data[11]
	m33 := m.data[15]

	n00 := n.data[0]
	n01 := n.data[4]
	n02 := n.data[8]
	n03 := n.data[12]
	n10 := n.data[1]
	n11 := n.data[5]
	n12 := n.data[9]
	n13 := n.data[13]
	n20 := n.data[2]
	n21 := n.data[6]
	n22 := n.data[10]
	n23 := n.data[14]
	n30 := n.data[3]
	n31 := n.data[7]
	n32 := n.data[11]
	n33 := n.data[15]

	m.data[0] = (m00 * n00) + (m01 * n10) + (m02 * n20) + (m03 * n30)
	m.data[4] = (m00 * n01) + (m01 * n11) + (m02 * n21) + (m03 * n31)
	m.data[8] = (m00 * n02) + (m01 * n12) + (m02 * n22) + (m03 * n32)
	m.data[12] = (m00 * n03) + (m01 * n13) + (m02 * n23) + (m03 * n33)
	m.data[1] = (m10 * n00) + (m11 * n10) + (m12 * n20) + (m13 * n30)
	m.data[5] = (m10 * n01) + (m11 * n11) + (m12 * n21) + (m13 * n31)
	m.data[9] = (m10 * n02) + (m11 * n12) + (m12 * n22) + (m13 * n32)
	m.data[13] = (m10 * n03) + (m11 * n13) + (m12 * n23) + (m13 * n33)
	m.data[2] = (m20 * n00) + (m21 * n10) + (m22 * n20) + (m23 * n30)
	m.data[6] = (m20 * n01) + (m21 * n11) + (m22 * n21) + (m23 * n31)
	m.data[10] = (m20 * n02) + (m21 * n12) + (m22 * n22) + (m23 * n32)
	m.data[14] = (m20 * n03) + (m21 * n13) + (m22 * n23) + (m23 * n33)
	m.data[3] = (m30 * n00) + (m31 * n10) + (m32 * n20) + (m33 * n30)
	m.data[7] = (m30 * n01) + (m31 * n11) + (m32 * n21) + (m33 * n31)
	m.data[11] = (m30 * n02) + (m31 * n12) + (m32 * n22) + (m33 * n32)
	m.data[15] = (m30 * n03) + (m31 * n13) + (m32 * n23) + (m33 * n33)
}

// DistanceSquared3 calculates the squared of the distance between two points.
func DistanceSquared3(x1, y1, z1, x2, y2, z2 float64) float64 {
	return LengthSquared3(x2-x1, y2-y1, z2-z1)
}

// Distance3 calculates the distance between two points.
func Distance3(x1, y1, z1, x2, y2, z2 float64) float64 {
	return Length3(x2-x1, y2-y1, z2-z1)
}

// Ortho3 reports if two vectors are orthogonal.
func Ortho3(x1, y1, z1, x2, y2, z2 float64) bool {
	return closeToZero(Dot3(x1, y1, z1, x2, y2, z2))
}

func closeToZero(f float64) bool {
	return math.Abs(f-0.0) < 0.000001
}

// Cross3 calculates the cross (vector) product of two vectors.
func Cross3(x1, y1, z1, x2, y2, z2 float64) (float64, float64, float64) {
	return y1*z2 - z1*y2, z1*x2 - x1*z2, x1*y2 - y1*x2
}

// Dot3 calculates the dot (scalar) product of two vectors.
func Dot3(x1, y1, z1, x2, y2, z2 float64) float64 {
	return x1*x2 + y1*y2 + z1*z2
}

// LengthSquared3 calculates the squared length of a vector.
func LengthSquared3(x, y, z float64) float64 {
	return x*x + y*y + z*z // dot3(x,y,z,x,y,z)
}

// Length3 calculates the length of a vector.
func Length3(x, y, z float64) float64 {
	return math.Sqrt(LengthSquared3(x, y, z))
}

// Normalize3 calculates a normalized copy of a vector.
func Normalize3(x, y, z float64) (float64, float64, float64) {
	length := Length3(x, y, z)
	if length == 0 {
		return x, y, z // ugh
	}
	return x / length, y / length, z / length
}

// SetRotationMatrix builds the rotation matrix.
// The rotation matrix will rotate a point from the null rotation direction to the direction specified by the given forward and up vectors.
//
// Null rotation vectors are:
// forward = 0 0 -1 // looking towards -Z
// up = 0 1 0       // up direction is +Y
// SetRotationMatrix(&rotation, 0, 0, -1, 0, 1, 0)
func SetRotationMatrix(rotationMatrix *Matrix4, forwardX, forwardY, forwardZ, upX, upY, upZ float64) {
	SetModelMatrix(rotationMatrix, forwardX, forwardY, forwardZ, upX, upY, upZ, 0, 0, 0)
}

// SetModelMatrix builds the model matrix.
// Model transformation represents objection location/orientation in world space.
// Model transformation is also known as "camera" transformation.
// Model transformation is the inverse of the view transformation.
// Common use is to compute object location/orientation into full transformation matrix.
//
// obj.coord. -> P*V*T*R*U*S -> clip coord -> divide by w -> NDC coord -> viewport transform -> window coord
// P*V*T*R*U*S = full transformation
// P = Perspective
// V = View (inverse of camera) built by setViewMatrix
// T*R = model transformation built by THIS setModelMatrix
// T = Translation
// R = Rotation
// U = Undo Model Local Rotation
// S = Scaling
//
// null model:
// forward = 0 0 -1    // looking towards -Z
// up = 0 1 0          // up direction is +Y
// translation = 0 0 0 // position at origin
// setModelMatrix(&rotation, 0, 0, -1, 0, 1, 0, 0, 0, 0)
func SetModelMatrix(modelMatrix *Matrix4, forwardX, forwardY, forwardZ, upX, upY, upZ, tX, tY, tZ float64) {
	rightX, rightY, rightZ := Normalize3(Cross3(forwardX, forwardY, forwardZ, upX, upY, upZ))

	rX := float32(rightX)
	rY := float32(rightY)
	rZ := float32(rightZ)

	uX := float32(upX)
	uY := float32(upY)
	uZ := float32(upZ)

	bX := -float32(forwardX)
	bY := -float32(forwardY)
	bZ := -float32(forwardZ)

	oX := float32(tX)
	oY := float32(tY)
	oZ := float32(tZ)

	modelMatrix.data[0] = rX
	modelMatrix.data[1] = rY
	modelMatrix.data[2] = rZ
	modelMatrix.data[3] = 0
	modelMatrix.data[4] = uX
	modelMatrix.data[5] = uY
	modelMatrix.data[6] = uZ
	modelMatrix.data[7] = 0
	modelMatrix.data[8] = bX
	modelMatrix.data[9] = bY
	modelMatrix.data[10] = bZ
	modelMatrix.data[11] = 0
	modelMatrix.data[12] = oX
	modelMatrix.data[13] = oY
	modelMatrix.data[14] = oZ
	modelMatrix.data[15] = 1
}

// SetViewMatrix builds the view matrix.
// View transformation represents camera inverted location/orientation in world space.
// View transformation moves all world objects in order to simulate a camera.
// View transformation is also known as "lookAt" transformation.
// View transformation is the inverse of the model transformation.
// Common use is to compute camera location/orientation into full transformation matrix.
//
// obj.coord. -> P*V*T*R*U*S -> clip coord -> divide by w -> NDC coord -> viewport transform -> window coord
// P*V*T*R*U*S = full transformation
// P = Perspective
// V = View (inverse of camera) built by THIS setViewMatrix
// T*R = model transformation built by setModelMatrix
// T = Translation
// R = Rotation
// U = Undo Model Local Rotation
// S = Scaling
//
// null view matrix:
// focus = 0 0 -1
// up    = 0 1 0
// pos   = 0 0 0
// setViewMatrix(&V, 0, 0, -1, 0, 1, 0, 0, 0, 0)
func SetViewMatrix(viewMatrix *Matrix4, focusX, focusY, focusZ, upX, upY, upZ, posX, posY, posZ float64) {

	backX, backY, backZ := Normalize3(posX-focusX, posY-focusY, posZ-focusZ)
	rightX, rightY, rightZ := Normalize3(Cross3(upX, upY, upZ, backX, backY, backZ))
	newUpX, newUpY, newUpZ := Normalize3(Cross3(backX, backY, backZ, rightX, rightY, rightZ))

	rotatedEyeX := -Dot3(rightX, rightY, rightZ, posX, posY, posZ)
	rotatedEyeY := -Dot3(newUpX, newUpY, newUpZ, posX, posY, posZ)
	rotatedEyeZ := -Dot3(backX, backY, backZ, posX, posY, posZ)

	rX := float32(rightX)
	rY := float32(rightY)
	rZ := float32(rightZ)

	uX := float32(newUpX)
	uY := float32(newUpY)
	uZ := float32(newUpZ)

	bX := float32(backX)
	bY := float32(backY)
	bZ := float32(backZ)

	eX := float32(rotatedEyeX)
	eY := float32(rotatedEyeY)
	eZ := float32(rotatedEyeZ)

	viewMatrix.data[0] = rX
	viewMatrix.data[1] = uX
	viewMatrix.data[2] = bX
	viewMatrix.data[3] = 0
	viewMatrix.data[4] = rY
	viewMatrix.data[5] = uY
	viewMatrix.data[6] = bY
	viewMatrix.data[7] = 0
	viewMatrix.data[8] = rZ
	viewMatrix.data[9] = uZ
	viewMatrix.data[10] = bZ
	viewMatrix.data[11] = 0
	viewMatrix.data[12] = eX
	viewMatrix.data[13] = eY
	viewMatrix.data[14] = eZ
	viewMatrix.data[15] = 1
}

// SetPerspectiveMatrix builds the perspective projection matrix.
func SetPerspectiveMatrix(perspectiveMatrix *Matrix4, fieldOfViewYRadians, aspectRatio, zNear, zFar float64) {

	f := math.Tan(math.Pi*0.5 - fieldOfViewYRadians*0.5) // = cotan(fieldOfViewYRadians/2)
	rangeInv := 1.0 / (zNear - zFar)

	d0 := float32(f / aspectRatio)
	d5 := float32(f)
	d10 := float32((zNear + zFar) * rangeInv)
	d14 := float32(zNear * zFar * rangeInv * 2.0)

	perspectiveMatrix.data[0] = d0
	perspectiveMatrix.data[1] = 0
	perspectiveMatrix.data[2] = 0
	perspectiveMatrix.data[3] = 0
	perspectiveMatrix.data[4] = 0
	perspectiveMatrix.data[5] = d5
	perspectiveMatrix.data[6] = 0
	perspectiveMatrix.data[7] = 0
	perspectiveMatrix.data[8] = 0
	perspectiveMatrix.data[9] = 0
	perspectiveMatrix.data[10] = d10
	perspectiveMatrix.data[11] = -1
	perspectiveMatrix.data[12] = 0
	perspectiveMatrix.data[13] = 0
	perspectiveMatrix.data[14] = d14
	perspectiveMatrix.data[15] = 0
}

/*
camera = includes both the perspective and view transforms

obj.coord. -> P*V*T*R*U*S -> clip coord -> divide by w -> NDC coord -> viewport transform -> window coord
P*V*T*R*U*S = full transformation
P = Perspective
V = View (inverse of camera) built by setViewMatrix
T*R = model transformation built by setModelMatrix
T = Translation
R = Rotation
U = Undo Model Local Rotation
S = Scaling
*/
func unproject(camera *Matrix4, viewportX, viewportWidth, viewportY, viewportHeight, pickX, pickY int, depth float64) (worldX, worldY, worldZ float64, err error) {

	// from screen coordinates to clip coordinates
	pX := (2.0 * float64(pickX-viewportX) / float64(viewportWidth)) - 1.0
	pY := (2.0 * float64(pickY-viewportY) / float64(viewportHeight)) - 1.0
	pZ := 2.0*depth - 1.0

	if pX < -1.0 || pX > 1.0 || pY < -1.0 || pY > 1.0 || pZ < -1.0 || pZ > 1.0 {
		err = errors.New("unproject: pick point outside unit cube")
		return
	}

	// invertedCamera: clip coord -> undo perspective -> undo view -> world coord
	var invertedCamera Matrix4
	invertedCamera.CopyInverseFrom(camera)
	vx, vy, vz, vw := invertedCamera.Transform(pX, pY, pZ, 1.0)
	if vw == 0.0 {
		err = errors.New("unproject: unprojected pick point with W=0")
		return
	}
	invW := 1.0 / vw
	worldX = vx * invW
	worldY = vy * invW
	worldZ = vz * invW

	return
}

// PickRay calculates points where pickX,pickY intersects near and far planes.
//
// camera = includes both the perspective and view transforms
// (camera: the func parameter)
//
// obj.coord. -> P*V*T*R*U*S -> clip coord -> divide by w -> NDC coord -> viewport transform -> window coord
// P*V*T*R*U*S = full transformation
// P = Perspective
// V = View (inverse of camera) built by setViewMatrix
// T*R = model transformation built by setModelMatrix
// T = Translation
// R = Rotation
// U = Undo Model Local Rotation
// S = Scaling
func PickRay(camera *Matrix4, viewportX, viewportWidth, viewportY, viewportHeight, pickX, pickY int) (nearX, nearY, nearZ, farX, farY, farZ float64, err error) {

	nearX, nearY, nearZ, err = unproject(camera, viewportX, viewportWidth, viewportY, viewportHeight, pickX, viewportHeight-pickY, 0.0)
	if err != nil {
		return
	}

	farX, farY, farZ, err = unproject(camera, viewportX, viewportWidth, viewportY, viewportHeight, pickX, viewportHeight-pickY, 1.0)

	return
}

// ViewportTransform simulates the viewport transform.
// ViewportTransform maps NDC coordinates to viewport coordinates.
// Input: viewport (viewportX, viewportWidth, viewportY, viewportHeight)
// Input: depthRange (depthNear, depthFar)
// Output: x,y,depth (x,y = viewport coord)
func ViewportTransform(viewportX, viewportWidth, viewportY, viewportHeight int, depthNear, depthFar, ndcX, ndcY, ndcZ float64) (int, int, float64) {
	return viewportTransform2(viewportX, viewportWidth, viewportY, viewportHeight, depthNear, depthFar, ndcX, ndcY, ndcZ)
}

func viewportTransform1(viewportX, viewportWidth, viewportY, viewportHeight int, depthNear, depthFar, ndcX, ndcY, ndcZ float64) (int, int, float64) {
	halfWidth := float64(viewportWidth) / 2.0
	halfHeight := float64(viewportHeight) / 2.0
	vx := roundToInt(ndcX*halfWidth+halfWidth) + viewportX
	vy := roundToInt(ndcY*halfHeight+halfHeight) + viewportY
	depth := (ndcZ*(depthFar-depthNear) + (depthFar + depthNear)) / 2.0

	return vx, viewportHeight - vy, depth
}

func viewportTransform2(viewportX, viewportWidth, viewportY, viewportHeight int, depthNear, depthFar, ndcX, ndcY, ndcZ float64) (int, int, float64) {
	halfWidth := .5 * float64(viewportWidth)
	halfHeight := .5 * float64(viewportHeight)
	vx := roundToInt(ndcX*halfWidth+halfWidth) + viewportX
	vy := roundToInt(ndcY*halfHeight+halfHeight) + viewportY
	depth := .5 * (ndcZ*(depthFar-depthNear) + (depthFar + depthNear))

	return vx, viewportHeight - vy, depth
}

func round(a float64) float64 {
	var r float64
	if a < 0 {
		r = math.Ceil(a - 0.5)
	} else {
		r = math.Floor(a + 0.5)
	}
	return r
}

func roundToInt(a float64) int {
	return int(round(a))
}

/*
SetOrthoMatrix builds matrix for orthographic projection.

near=-1 far=1 -> flip Z (this is the usual ortho projection)
near=1 far=-1 -> keep Z

SetOrthoMatrix(m,-1,1,-1,1,-1,1): flip Z (this is the usual ortho projection)
SetOrthoMatrix(m,-1,1,-1,1,1,-1): identity
*/
func SetOrthoMatrix(orthoMatrix *Matrix4, left, right, bottom, top, near, far float64) {
	lr := 1.0 / (left - right)
	bt := 1.0 / (bottom - top)
	nf := 1.0 / (near - far)
	orthoMatrix.data[0] = float32(-2.0 * lr)
	orthoMatrix.data[1] = 0
	orthoMatrix.data[2] = 0
	orthoMatrix.data[3] = 0
	orthoMatrix.data[4] = 0
	orthoMatrix.data[5] = float32(-2.0 * bt)
	orthoMatrix.data[6] = 0
	orthoMatrix.data[7] = 0
	orthoMatrix.data[8] = 0
	orthoMatrix.data[9] = 0
	orthoMatrix.data[10] = float32(2.0 * nf)
	orthoMatrix.data[11] = 0
	orthoMatrix.data[12] = float32((left + right) * lr)
	orthoMatrix.data[13] = float32((top + bottom) * bt)
	orthoMatrix.data[14] = float32((far + near) * nf)
	orthoMatrix.data[15] = 1
}
