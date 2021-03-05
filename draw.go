package main

import (
	"fmt"
	"time"
)

var draw string = `
           +                 +
          / \               / \
         /   \             /   \
        /     \           /     \
       /       \         /       \
      +         \       /         +
       \         \     /         /  
        \         \   /         /  
         \         \ /         /  
          \         0         /
           \       / \       /
            \     /   \     /
             \   /     \   /
              \ /       \ /
               +         +
`

func StartWithDrawing() {
	a := MakeThreeD(zero, "a")
	b := MakeThreeD(zero, "b")
	go a.StartDraw()
	go b.StartDraw()
	for {
		time.Sleep(time.Second * 1)
	}
}
func DrawWithParams(ab string) {
	draw := `
           +                 +
          / \               / \
         /   \             /   \
        /     \           /     \
       /       \         /       \
      +         \       /         +
       \         \     /         /  
        \         \   /         /  
         \         \ /         /  
          \         %s         /
           \       / \       /
            \     /   \     /
             \   /     \   /
              \ /       \ /
               +         +
`
	fmt.Printf("\n\n\n%s\n\n\n", fmt.Sprintf(draw, ab))
}
func (t *Tron) PingDraw(from3D *ThreeD) []int {
	list := []int{}
	if t.Name == "zero" {
		//fmt.Printf("The SHARED ZERO hit by %s\n", from3D.Name)
		list = cacheTimes
		if from3D.Name == "a" {
			DrawWithParams("a")
			return list
		} else {
			DrawWithParams("b")
			return list
		}
	} else {
		//fmt.Printf("%s is at position %s\n", from3D.Name, t.Name)
	}
	fmt.Printf("\n\n\n%s\n\n\n", draw)
	return list
}
func (td *ThreeD) StartDraw() {
	for {
		lock.Lock()
		listOfABS[td.Name]++
		times := td.Zero.PingDraw(td)
		for i, _ := range td.Trons {
			//fmt.Printf("Delay %d milliseconds.\n", times[i])
			time.Sleep(time.Millisecond * time.Duration(times[i]))
			//tron.PingDraw(td)
			listOfABS[td.Name]++
		}
		if listOfABS["a"] >= 4 && listOfABS["b"] >= 4 {
			cacheTimes = MakeTimes(false)
			listOfABS = map[string]int{}
		}
		lock.Unlock()
	}
}
