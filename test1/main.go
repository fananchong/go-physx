package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
	"unsafe"

	gophysx "github.com/fananchong/go-physx"
)

/*
#cgo LDFLAGS: -L./ -lPProfGo
#include "PProfGo.h"
#include "stdlib.h"
*/
import "C"

const (
	DEFAULT_TEST_COUNT = 100
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.NewSource(time.Now().UnixNano())
	var wg sync.WaitGroup

	HeapProfilerStart("test")

	HeapProfilerDump("mem_init_sdk")

	for i := 0; i < DEFAULT_TEST_COUNT; i++ {
		wg.Add(1)
		go func() {
			scene, err := gophysx.NewScene("")
			if err != nil {
				panic(err)
			}
			var actors []uint64
			for j := 0; j < 1000; j++ {
				x := rand.Float32() * 100
				y := rand.Float32() * 100
				z := rand.Float32() * 100
				r := rand.Float32() * 10
				actor1 := scene.CreateSphereDynamic(gophysx.Vector3{x, y, z}, r)
				actors = append(actors, actor1)
				actor2 := scene.CreateBoxStatic(gophysx.Vector3{x + 10, y + 10, z + 10}, gophysx.Vector3{1, 1, 1})
				actors = append(actors, actor2)
			}

			var temphold float32
			for j := 0; j < 500; j++ {
				for _, actor := range actors {
					x, y, z := scene.GetGlobalPostion(actor)
					temphold = temphold + x + y + z
				}
				scene.Update(0.016)
			}

			scene.Release()
			scene = nil
			fmt.Println(".", temphold)
			wg.Done()
		}()
	}

	wg.Wait()

	HeapProfilerDump("mem_end_scene")
	gophysx.ReleasePhysxSDK()
	HeapProfilerDump("mem_release_sdk")

	HeapProfilerStop()
}

func HeapProfilerStart(f string) {
	spath := C.CString(f)
	defer C.free(unsafe.Pointer(spath))
	C._HeapProfilerStart(spath)
}

func HeapProfilerDump(s string) {
	temps := C.CString(s)
	defer C.free(unsafe.Pointer(temps))
	C._HeapProfilerDump(temps)
}

func HeapProfilerStop() {
	C._HeapProfilerStop()
}

/*
func HeapProfilerDump(filename string) {
	f, err := os.Create(filename + ".prof")
	if err != nil {
		log.Fatal(err)
		return
	}
	pprof.WriteHeapProfile(f)
}
*/

