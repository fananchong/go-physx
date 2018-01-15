package gophysx

/*
#cgo LDFLAGS: -L./ -lPhysxWrapGo
#include "PhysxWrapGo.h"
*/
import "C"
import "errors"

var bInitSDK int

func init() {
	bInitSDK = int(C.InitPhysxSDK())
	if bInitSDK == 0 {
		panic("[FATAL] CANT INIT PHYSX SDK!")
	}
}

var ErrNeedInitSDK = errors.New("physx sdk is not init.")
var ErrCreateSceneFail = errors.New("create scene fail.")

type PxScene struct {
	c C.PxScene
}

func NewScene(path string) (*PxScene, error) {
	if bInitSDK == 0 {
		return nil, ErrNeedInitSDK
	}
	this := &PxScene{}
	this.c = C.CreateScene(C.CString(path))
	if this.c == nil {
		return nil, ErrCreateSceneFail
	}
	return this, nil
}

func (this *PxScene) Release() {
	if this.c != nil {
		C.DestroyScene(this.c)
		this.c = nil
	}
}

func (this *PxScene) Update(elapsedTime float32) {
	if this.c != nil {
		C.UpdateScene(this.c, C.float(elapsedTime))
	}
}

func (this *PxScene) CreatePlane(yAxis float32) (id uint64) {
	if this.c != nil {
		id = uint64(C.CreatePlane(this.c, C.float(yAxis)))
	}
	return
}

func (this *PxScene) CreateBoxDynamic(pos Vector3, halfExtents Vector3) (id uint64) {
	if this.c != nil {
		cPos := C.Vector3{}
		cPos.X = C.float(pos.X)
		cPos.Y = C.float(pos.Y)
		cPos.Z = C.float(pos.Z)
		cHalfExtents := C.Vector3{}
		cHalfExtents.X = C.float(halfExtents.X)
		cHalfExtents.Y = C.float(halfExtents.Y)
		cHalfExtents.Z = C.float(halfExtents.Z)
		id = uint64(C.CreateBoxDynamic(this.c, &cPos, &cHalfExtents))
	}
	return
}

func (this *PxScene) CreateBoxKinematic(pos Vector3, halfExtents Vector3) (id uint64) {
	if this.c != nil {
		cPos := C.Vector3{}
		cPos.X = C.float(pos.X)
		cPos.Y = C.float(pos.Y)
		cPos.Z = C.float(pos.Z)
		cHalfExtents := C.Vector3{}
		cHalfExtents.X = C.float(halfExtents.X)
		cHalfExtents.Y = C.float(halfExtents.Y)
		cHalfExtents.Z = C.float(halfExtents.Z)
		id = uint64(C.CreateBoxKinematic(this.c, &cPos, &cHalfExtents))
	}
	return
}

func (this *PxScene) CreateBoxStatic(pos Vector3, halfExtents Vector3) (id uint64) {
	if this.c != nil {
		cPos := C.Vector3{}
		cPos.X = C.float(pos.X)
		cPos.Y = C.float(pos.Y)
		cPos.Z = C.float(pos.Z)
		cHalfExtents := C.Vector3{}
		cHalfExtents.X = C.float(halfExtents.X)
		cHalfExtents.Y = C.float(halfExtents.Y)
		cHalfExtents.Z = C.float(halfExtents.Z)
		id = uint64(C.CreateBoxStatic(this.c, &cPos, &cHalfExtents))
	}
	return
}

func (this *PxScene) CreateSphereDynamic(pos Vector3, radius float32) (id uint64) {
	if this.c != nil {
		cPos := C.Vector3{}
		cPos.X = C.float(pos.X)
		cPos.Y = C.float(pos.Y)
		cPos.Z = C.float(pos.Z)
		id = uint64(C.CreateSphereDynamic(this.c, &cPos, C.float(radius)))
	}
	return
}

func (this *PxScene) CreateSphereKinematic(pos Vector3, radius float32) (id uint64) {
	if this.c != nil {
		cPos := C.Vector3{}
		cPos.X = C.float(pos.X)
		cPos.Y = C.float(pos.Y)
		cPos.Z = C.float(pos.Z)
		id = uint64(C.CreateSphereKinematic(this.c, &cPos, C.float(radius)))
	}
	return
}

func (this *PxScene) CreateSphereStatic(pos Vector3, radius float32) (id uint64) {
	if this.c != nil {
		cPos := C.Vector3{}
		cPos.X = C.float(pos.X)
		cPos.Y = C.float(pos.Y)
		cPos.Z = C.float(pos.Z)
		id = uint64(C.CreateSphereStatic(this.c, &cPos, C.float(radius)))
	}
	return
}

func (this *PxScene) CreateCapsuleDynamic(pos Vector3, radius, halfHeight float32) (id uint64) {
	if this.c != nil {
		cPos := C.Vector3{}
		cPos.X = C.float(pos.X)
		cPos.Y = C.float(pos.Y)
		cPos.Z = C.float(pos.Z)
		id = uint64(C.CreateCapsuleDynamic(this.c, &cPos, C.float(radius), C.float(halfHeight)))
	}
	return
}

func (this *PxScene) CreateCapsuleKinematic(pos Vector3, radius, halfHeight float32) (id uint64) {
	if this.c != nil {
		cPos := C.Vector3{}
		cPos.X = C.float(pos.X)
		cPos.Y = C.float(pos.Y)
		cPos.Z = C.float(pos.Z)
		id = uint64(C.CreateCapsuleKinematic(this.c, &cPos, C.float(radius), C.float(halfHeight)))
	}
	return
}

func (this *PxScene) CreateCapsuleStatic(pos Vector3, radius, halfHeight float32) (id uint64) {
	if this.c != nil {
		cPos := C.Vector3{}
		cPos.X = C.float(pos.X)
		cPos.Y = C.float(pos.Y)
		cPos.Z = C.float(pos.Z)
		id = uint64(C.CreateCapsuleStatic(this.c, &cPos, C.float(radius), C.float(halfHeight)))
	}
	return
}

func (this *PxScene) RemoveActor(id uint64) {
	if this.c != nil {
		C.RemoveActor(this.c, C.ulonglong(id))
	}
}

func (this *PxScene) SetLinearVelocity(id uint64, velocity Vector3) {
	if this.c != nil {
		cVelocity := C.Vector3{}
		cVelocity.X = C.float(velocity.X)
		cVelocity.Y = C.float(velocity.Y)
		cVelocity.Z = C.float(velocity.Z)
		C.SetLinearVelocity(this.c, C.ulonglong(id), &cVelocity)
	}
}

func (this *PxScene) AddForce(id uint64, force Vector3) {
	if this.c != nil {
		cForce := C.Vector3{}
		cForce.X = C.float(force.X)
		cForce.Y = C.float(force.Y)
		cForce.Z = C.float(force.Z)
		C.AddForce(this.c, C.ulonglong(id), &cForce)
	}
}

func (this *PxScene) ClearForce(id uint64) {
	if this.c != nil {
		C.ClearForce(this.c, C.ulonglong(id))
	}
}

func (this *PxScene) GetGlobalPostion(id uint64) (outPostion Vector3) {
	if this.c != nil {
		cPos := C.Vector3{}
		C.GetGlobalPostion(this.c, C.ulonglong(id), &cPos)
		outPostion.X = float32(cPos.X)
		outPostion.Y = float32(cPos.Y)
		outPostion.Z = float32(cPos.Z)
	}
	return
}

func (this *PxScene) GetGlobalRotate(id uint64) (outRotate Quat) {
	if this.c != nil {
		cRotate := C.Quat{}
		C.GetGlobalRotate(this.c, C.ulonglong(id), &cRotate)
		outRotate.X = float32(cRotate.X)
		outRotate.Y = float32(cRotate.Y)
		outRotate.Z = float32(cRotate.Z)
		outRotate.W = float32(cRotate.W)
	}
	return
}

func (this *PxScene) SetGlobalPostion(id uint64, pos Vector3) {
	if this.c != nil {
		cPos := C.Vector3{}
		cPos.X = C.float(pos.X)
		cPos.Y = C.float(pos.Y)
		cPos.Z = C.float(pos.Z)
		C.SetGlobalPostion(this.c, C.ulonglong(id), &cPos)
	}
}

func (this *PxScene) SetGlobalRotate(id uint64, rotate Quat) {
	if this.c != nil {
		cRotate := C.Quat{}
		cRotate.X = C.float(rotate.X)
		cRotate.Y = C.float(rotate.Y)
		cRotate.Z = C.float(rotate.Z)
		cRotate.Z = C.float(rotate.W)
		C.SetGlobalRotate(this.c, C.ulonglong(id), &cRotate)
	}
}
