package gophysx

/*
#cgo LDFLAGS: -L./ -lPhysxWrapGo
#include "PhysxWrapGo.h"
#include "stdlib.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

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
	c unsafe.Pointer
}

func NewScene(path string) (*PxScene, error) {
	spath := C.CString(path)
	defer C.free(unsafe.Pointer(spath))
	if bInitSDK == 0 {
		return nil, ErrNeedInitSDK
	}
	this := &PxScene{}
	this.c = C.CreateScene(spath)
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
		id = uint64(C.CreateBoxDynamic(this.c, C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(halfExtents.X), C.float(halfExtents.Y), C.float(halfExtents.Z)))
	}
	return
}

func (this *PxScene) CreateBoxKinematic(pos Vector3, halfExtents Vector3) (id uint64) {
	if this.c != nil {
		id = uint64(C.CreateBoxKinematic(this.c, C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(halfExtents.X), C.float(halfExtents.Y), C.float(halfExtents.Z)))
	}
	return
}

func (this *PxScene) CreateBoxStatic(pos Vector3, halfExtents Vector3) (id uint64) {
	if this.c != nil {
		id = uint64(C.CreateBoxStatic(this.c, C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(halfExtents.X), C.float(halfExtents.Y), C.float(halfExtents.Z)))
	}
	return
}

func (this *PxScene) CreateSphereDynamic(pos Vector3, radius float32) (id uint64) {
	if this.c != nil {
		id = uint64(C.CreateSphereDynamic(this.c, C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(radius)))
	}
	return
}

func (this *PxScene) CreateSphereKinematic(pos Vector3, radius float32) (id uint64) {
	if this.c != nil {
		id = uint64(C.CreateSphereKinematic(this.c, C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(radius)))
	}
	return
}

func (this *PxScene) CreateSphereStatic(pos Vector3, radius float32) (id uint64) {
	if this.c != nil {
		id = uint64(C.CreateSphereStatic(this.c, C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(radius)))
	}
	return
}

func (this *PxScene) CreateCapsuleDynamic(pos Vector3, radius, halfHeight float32) (id uint64) {
	if this.c != nil {
		id = uint64(C.CreateCapsuleDynamic(this.c, C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(radius), C.float(halfHeight)))
	}
	return
}

func (this *PxScene) CreateCapsuleKinematic(pos Vector3, radius, halfHeight float32) (id uint64) {
	if this.c != nil {
		id = uint64(C.CreateCapsuleKinematic(this.c, C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(radius), C.float(halfHeight)))
	}
	return
}

func (this *PxScene) CreateCapsuleStatic(pos Vector3, radius, halfHeight float32) (id uint64) {
	if this.c != nil {
		id = uint64(C.CreateCapsuleStatic(this.c, C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(radius), C.float(halfHeight)))
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
		C.SetLinearVelocity(this.c, C.ulonglong(id), C.float(velocity.X), C.float(velocity.Y), C.float(velocity.Z))
	}
}

func (this *PxScene) AddForce(id uint64, force Vector3) {
	if this.c != nil {
		C.AddForce(this.c, C.ulonglong(id), C.float(force.X), C.float(force.Y), C.float(force.Z))
	}
}

func (this *PxScene) ClearForce(id uint64) {
	if this.c != nil {
		C.ClearForce(this.c, C.ulonglong(id))
	}
}

func (this *PxScene) GetGlobalPostion(id uint64) (outPostionX, outPostionY, outPostionZ float32) {
	if this.c != nil {
		C.GetGlobalPostion(this.c, C.ulonglong(id), unsafe.Pointer(&outPostionX), unsafe.Pointer(&outPostionY), unsafe.Pointer(&outPostionZ))
	}
	return
}

func (this *PxScene) GetGlobalRotate(id uint64) (outRotateX, outRotateY, outRotateZ, outRotateW float32) {
	if this.c != nil {
		C.GetGlobalRotate(this.c, C.ulonglong(id), unsafe.Pointer(&outRotateX), unsafe.Pointer(&outRotateY), unsafe.Pointer(&outRotateZ), unsafe.Pointer(&outRotateW))
	}
	return
}

func (this *PxScene) SetGlobalPostion(id uint64, pos Vector3) {
	if this.c != nil {
		C.SetGlobalPostion(this.c, C.ulonglong(id), C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
	}
}

func (this *PxScene) SetGlobalRotate(id uint64, rotate Quat) {
	if this.c != nil {
		C.SetGlobalRotate(this.c, C.ulonglong(id), C.float(rotate.X), C.float(rotate.Y), C.float(rotate.Z), C.float(rotate.W))
	}
}

func (this *PxScene) IsStaticObj(id uint64) (ok bool) {
	if this.c != nil {
		ok = (C.IsStaticObj(this.c, C.ulonglong(id)) != 0)
	}
	return
}

func (this *PxScene) IsDynamicObj(id uint64) (ok bool) {
	if this.c != nil {
		ok = (C.IsDynamicObj(this.c, C.ulonglong(id)) != 0)
	}
	return
}

func (this *PxScene) SetCurrentMaterial(staticFriction, dynamicFriction, restitution float32) {
	if this.c != nil {
		C.SetCurrentMaterial(this.c, C.float(staticFriction), C.float(dynamicFriction), C.float(restitution))
	}
}

func (this *PxScene) SetCurrentAngularDamping(value float32) {
	if this.c != nil {
		C.SetCurrentAngularDamping(this.c, C.float(value))
	}
}
