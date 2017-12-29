package gophysx

/*
#cgo LDFLAGS: -L./ -lPhysxWrapGo
#include "PhysxWrapGo.h"
*/
import "C"
import "errors"

var bInitSDK bool

func init() {
	bInitSDK := C.InitPhysxSDK()
	if !bInitSDK {
		panic("[FATAL] CANT INIT PHYSX SDK!")
	}
}

var ErrNeedInitSDK = errors.New("physx sdk is not init.")
var ErrCreateSceneFail = errors.New("create scene fail.")

type PxScene struct {
	c C.PxScene
}

func NewScene(path string) (*PxScene, error) {
	if !bInitSDK {
		return nil, ErrNeedInitSDK
	}
	this := &PxScene{}
	this.c = C.CreateScene(C.CString(path))
	if this.c == nil {
		return nil, ErrCreateSceneFail
	}
	return this, nil
}

//func (this *PxScene) DestroyScene() {
//	if this.c != nil {
//		C.DestroyScene(this.c)
//		this.c = nil
//	}
//}
//
//func (this *PxScene) Update(elapsedTime float) {
//	if this.c != nil {
//		C.UpdateScene(this.c, C.float(elapsedTime))
//	}
//}
