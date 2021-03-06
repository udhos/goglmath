[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/udhos/goglmath/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/udhos/goglmath?status.svg)](http://godoc.org/github.com/udhos/goglmath)
[![Go Report Card](https://goreportcard.com/badge/github.com/udhos/goglmath)](https://goreportcard.com/report/github.com/udhos/goglmath)

# goglmath
goglmath - Lightweight pure Go 3D math package providing essential matrix/vector operations for GL graphics applications.

Full Transformation Stack
=========================

    obj.coord. -> P*V*T*R*U*S*o -> clip coord  -> divide by w -> NDC coord -> viewport transform      -> win.coord+depth
                    ---------      -----------                   ---------    -----------------------    ---------------
                    "MV"           gl_Position                   vec3         Viewport()+DepthRange()    x,y,depth
                                   vec4                          -1..1

    o           = obj.coord
    P*V*T*R*U*S = full transformation matrix
    P           = Perspective
    V           = View (inverse of camera) built by setViewMatrix
    T*R         = Model transformation built by setModelMatrix
    T           = Translation
    R           = Rotation
    U           = Undo Model Local Rotation
    S           = Scaling

    Typical vertex shader: gl_Position = u_P * u_MV * vec4(a_Position, 1.0);
    a_Position  = obj.coord
    u_MV        = V*T*R*U*S
    u_P         = P
    gl_Position = clip coord

    Viewport transform:
    - After clipping and division by w, NDC coordinates range from -1 to 1
    - programmed with: Viewport()+DepthRange()
    - input: NDC coord (vec3)
    - output: win.coord (x,y) + depth

--xx--
