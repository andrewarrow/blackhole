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
	a := MakeThreeD(zero, "a")
	b := MakeThreeD(zero, "b")
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

func DrawWithParams(ab string, num int) {
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
	if ab == "a" {
		thing.Zero = "a"
		if num == 0 {
			//fmt.Printf("%s\n", fmt.Sprintf(draw, "~",
			//"+", "+", "+", "+", "a", "(", ")", "+", "+"))
		} else if num == 1 || (num == 2 && rand.Intn(20) > 5) {
			thing.OneLeft = "1"
			//fmt.Printf("%s\n", fmt.Sprintf(draw, "~",
			//"+", "1", "+", "+", "a", "(", ")", "+", "+"))
		} else if num == 2 || (num == 1 && rand.Intn(20) <= 5) {
			thing.TwoLeft = "2"
			//fmt.Printf("%s\n", fmt.Sprintf(draw, "~",
			//"+", "+", "+", "2", "a", "(", ")", "+", "+"))
		} else if num >= 3 {
			thing.ThreeLeft = "3"
			//fmt.Printf("%s\n", fmt.Sprintf(draw, "~",
			//"+", "+", "+", "+", "a", "(", ")", "+", "3"))
		}
	} else {
		thing.Zero = "b"
		if num == 0 {
			//	fmt.Printf("%s\n", fmt.Sprintf(draw, "~",
			//	"+", "+", "+", "+", "b", "(", ")", "+", "+"))
		} else if num == 1 || (num == 2 && rand.Intn(20) > 5) {
			thing.OneRight = "1"
			//fmt.Printf("%s\n", fmt.Sprintf(draw, "~",
			//"1", "+", "+", "+", "b", "(", ")", "+", "+"))
		} else if num == 2 || (num == 1 && rand.Intn(20) <= 5) {
			thing.TwoRight = "2"
			//fmt.Printf("%s\n", fmt.Sprintf(draw, "~",
			//"+", "+", "2", "+", "b", "(", ")", "+", "+"))
		} else if num >= 3 {
			thing.ThreeRight = "3"
			//fmt.Printf("%s\n", fmt.Sprintf(draw, "~",
			//"+", "+", "+", "+", "b", "(", ")", "3", "+"))
		}
	}

	err = tmpl.Execute(os.Stdout, thing)
	fmt.Printf("completedRevs: %d, %f", completedRevs, float64(totalSums)/float64(totalCount))
}
func (t *Tron) PingDraw(from3D *ThreeD) []int {
	list := []int{}
	if t.Name == "zero" {
		//fmt.Printf("The SHARED ZERO hit by %s\n", from3D.Name)
		list = cacheTimes
		if from3D.Name == "a" {
			DrawWithParams("a", 0)
		} else {
			DrawWithParams("b", 0)
		}
	} else {
		//fmt.Printf("%s is at position %s\n", from3D.Name, t.Name)
		if from3D.Name == "a" {
			DrawWithParams("a", listOfABS[from3D.Name]+1)
		} else {
			DrawWithParams("b", listOfABS[from3D.Name]+1)
		}
	}
	return list
}
func (td *ThreeD) StartDraw() {
	for {
		lock.Lock()
		listOfABS[td.Name]++
		time.Sleep(time.Millisecond * 100)
		times := td.Zero.PingDraw(td)
		for i, tron := range td.Trons {
			//fmt.Printf("Delay %d milliseconds.\n", times[i])
			time.Sleep(time.Millisecond * time.Duration(times[i]))
			tron.PingDraw(td)
			listOfABS[td.Name]++
		}
		if listOfABS["a"] >= 4 && listOfABS["b"] >= 4 {
			cacheTimes = MakeTimes(false)
			listOfABS = map[string]int{}
			completedRevs++
		}
		lock.Unlock()
	}
}
