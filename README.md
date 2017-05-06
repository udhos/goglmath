# goglmath
goglmath - Lightweight pure Go 3D math package providing essential matrix/vector operations for GL graphics applications.

Full Transformation Stack
=========================

    obj.coord. -> P*V*T*R*U*S*o -> clip coord  -> divide by w -> NDC coord -> viewport transform -> window coord
                    ---------      -----------
                    "MV"           gl_Position

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

--xx--
