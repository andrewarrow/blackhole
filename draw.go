package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"os"
	"time"
)

var completedRevs int

func StartWithDrawing() {
	a := MakeThreeD(zeroFor3, "a")
	b := MakeThreeD(zeroFor3, "b")
	c := MakeFourD(zeroFor4, "c")
	go c.StartDraw()
	go a.StartDraw()
	go b.StartDraw()
	for {
		time.Sleep(time.Second * 1)
	}
}

type TemplateThing struct {
	Top        string
	OneLeft    string
	OneRight   string
	TwoLeft    string
	TwoRight   string
	ThreeLeft  string
	ThreeRight string
	Zero       string
	CloseLeft  string
	CloseRight string
}

type DrawParams struct {
	Ab  string
	Num int
}

func DrawWithParams(dp DrawParams) {
	draw := `
                    {{ .Top }}
                   / \ 
           {{ .OneLeft }}      /   \      {{ .OneRight }}
          / \    /     \    / \
         /   \  /       \  /   \
        /     \/         \/     \
       /      /\         /\      \
      {{ .TwoLeft }}      /  \       /  \      {{ .TwoRight }}
       \    /    \     /    \    /  
        \  /      \   /      \  /  
         \/        \ /        \/         
         /\         {{ .Zero }}         /\
        /  \       / \       /  \
       {{ .CloseLeft }}    \     /   \     /    {{ .CloseRight }}     
             \   /     \   /      
              \ /       \ /
               {{ .ThreeLeft }}        {{ .ThreeRight }}


`

	tmpl, err := template.New("test").Parse(draw)
	if err != nil {
		panic(err)
	}
	thing := TemplateThing{}
	thing.Top = "~"
	thing.OneLeft = "."
	thing.OneRight = "."
	thing.TwoLeft = "."
	thing.TwoRight = "."
	thing.ThreeLeft = "."
	thing.ThreeRight = "."
	thing.CloseLeft = "("
	thing.CloseRight = ")"
	if err != nil {
		panic(err)
	}
	if dp.Ab == "a" {
		thing.Zero = "a"
		if dp.Num == 0 {
		} else if dp.Num == 1 || (dp.Num == 2 && rand.Intn(20) > 5) {
			thing.OneLeft = "1"
		} else if dp.Num == 2 || (dp.Num == 1 && rand.Intn(20) <= 5) {
			thing.TwoLeft = "2"
		} else if dp.Num >= 3 {
			thing.ThreeLeft = "3"
		}
	} else {
		thing.Zero = "b"
		if dp.Num == 0 {
		} else if dp.Num == 1 || (dp.Num == 2 && rand.Intn(20) > 5) {
			thing.OneRight = "1"
		} else if dp.Num == 2 || (dp.Num == 1 && rand.Intn(20) <= 5) {
			thing.TwoRight = "2"
		} else if dp.Num >= 3 {
			thing.ThreeRight = "3"
		}
	}

	tmpl.Execute(os.Stdout, thing)
	fmt.Printf("completedRevs: %d, %f", completedRevs, float64(totalSums)/float64(totalCount))
}
func (t *Tron) PingDraw(from3D *ThreeD) []int {
	list := []int{}
	if t.Name == "zero" {
		//fmt.Printf("The SHARED ZERO hit by %s\n", from3D.Name)
		list = cacheTimes
		if from3D.Name == "a" {
			DrawWithParams(DrawParams{"a", 0})
		} else {
			DrawWithParams(DrawParams{"b", 0})
		}
	} else {
		//fmt.Printf("%s is at position %s\n", from3D.Name, t.Name)
		if from3D.Name == "a" {
			DrawWithParams(DrawParams{"a", listOfABS3[from3D.Name] + 1})
		} else {
			DrawWithParams(DrawParams{"b", listOfABS3[from3D.Name] + 1})
		}
	}
	return list
}
func (td *ThreeD) StartDraw() {
	for {
		lock.Lock()
		listOfABS3[td.Name]++
		time.Sleep(time.Millisecond * 100)
		times := td.Zero.PingDraw(td)
		for i, tron := range td.Trons {
			//fmt.Printf("Delay %d milliseconds.\n", times[i])
			time.Sleep(time.Millisecond * time.Duration(times[i]))
			tron.PingDraw(td)
			listOfABS3[td.Name]++
		}
		if listOfABS3["a"] >= 4 && listOfABS3["b"] >= 4 {
			cacheTimes = MakeTimes(false)
			listOfABS3 = map[string]int{}
			completedRevs++
		}
		lock.Unlock()
	}
}
