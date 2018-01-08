package main

import (
	"fmt"
	"time"

	gophysx "github.com/fananchong/go-physx"
)

func main() {

	{
		for i := 0; i < 10; i++ {
			scene, err := gophysx.NewScene("pxscene")
			if err != nil {
				panic(err)
			}
			scene.Release()
		}
	}

	scene, err := gophysx.NewScene("")
	if err != nil {
		panic(err)
	}

	scene.CreatePlane(0)
	actor := scene.CreateSphereDynamic(gophysx.Vector3{0, 25, 0}, 25)
	scene.SetLinearVelocity(actor, gophysx.Vector3{0, 0, 1})

	preTime := time.Now().UnixNano()
	t := time.NewTicker(33 * time.Millisecond)
	for {
		select {
		case <-t.C:
			{
				nowTime := time.Now().UnixNano()
				diff := float32(nowTime-preTime) / float32(time.Second)
				preTime = nowTime
				scene.Update(diff)
				pos := scene.GetGlobalPostion(actor)
				fmt.Println("diff =", diff)
				fmt.Printf("(x, y, z) = (%f,%f,%f)\n", pos.X, pos.Y, pos.Z)
			}
		}
	}
	scene.Release()
}
