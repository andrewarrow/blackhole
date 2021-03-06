package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ThreeD struct {
	Zero  *Tron
	Name  string
	Trons []*Tron
}

func MakeThreeD(zero *Tron, name string) *ThreeD {
	td := ThreeD{}
	td.Zero = zero
	td.Name = name
	td.Trons = []*Tron{}

	t1 := MakeTron("t1")
	t2 := MakeTron("t2")
	t3 := MakeTron("t3")
	td.Trons = append(td.Trons, t1)
	td.Trons = append(td.Trons, t2)
	td.Trons = append(td.Trons, t3)
	return &td
}

func MakeTimes3(verbose bool) []int {
	T := []int{}
	T = append(T, rand.Intn(500)+50)
	T = append(T, rand.Intn(500)+50)
	T = append(T, rand.Intn(500)+50)
	sum := 0
	for _, t := range T {
		sum += t
	}
	totalSums3 += sum
	totalCount3++
	if verbose {
		fmt.Printf("\nNext cycle will take %d milliseconds (Random between 0 and 1500).\n\n", sum)
	}
	return T
}

func (td *ThreeD) Start() {
	for {
		lock.Lock()
		listOfABS3[td.Name]++
		times := td.Zero.Ping(td)
		for i, tron := range td.Trons {
			fmt.Printf("Delay %d milliseconds.\n", times[i])
			time.Sleep(time.Millisecond * time.Duration(times[i]))
			tron.Ping(td)
			listOfABS3[td.Name]++
		}
		if listOfABS3["a"] >= 4 && listOfABS3["b"] >= 4 {
			cacheTimes3 = MakeTimes3(true)
			listOfABS3 = map[string]int{}
		}
		lock.Unlock()
	}
}

func Start() {
	a := MakeThreeD(zeroFor3, "a")
	b := MakeThreeD(zeroFor3, "b")
	go a.Start()
	go b.Start()

	for {
		time.Sleep(time.Second * 1)
	}
}
