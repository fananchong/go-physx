package main

import (
	"time"

	gophysx "github.com/fananchong/go-physx"
)

func main() {

	{
		for i := 0; i < 10; i++ {
			_, err := gophysx.NewScene("pxscene")
			if err != nil {
				panic(err)
			}
		}
	}

	scene, err := gophysx.NewScene("pxscene")
	if err != nil {
		panic(err)
	}

	actor := scene.CreateSphereDynamic(gophysx.Vector3{10, 25, 10}, 25)
	scene.SetLinearVelocity(actor, gophysx.Vector3{0, 0, 1})

	preTime := time.Now().UnixNano()
	t := time.NewTicker(33 * time.Millisecond)
	for {
		select {
		case <-t.C:
			{
				nowTime := time.Now().UnixNano()
				diff := float32(nowTime-preTime) / float32(time.Millisecond)
				scene.Update(diff)
			}
		}
	}
}
