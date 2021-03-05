package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var listOfABS map[string]int = map[string]int{}
var cacheTimes []int = MakeTimes()
var lock *sync.Mutex

type Tron struct {
	Name string
}

func MakeTron(name string) *Tron {
	t := Tron{}
	t.Name = name
	return &t
}

func (t *Tron) Ping(from3D *ThreeD) []int {
	list := []int{}
	if t.Name == "zero" {
		fmt.Printf("The SHARED ZERO hit by %s\n", from3D.Name)
		list = cacheTimes
	} else {
		fmt.Printf("%s is at position %s\n", from3D.Name, t.Name)
	}
	return list
}

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

func MakeTimes() []int {
	T := []int{}
	T = append(T, rand.Intn(500))
	T = append(T, rand.Intn(500))
	T = append(T, rand.Intn(500))
	sum := 0
	for _, t := range T {
		sum += t
	}
	fmt.Printf("\nNext cycle will take %d milliseconds (Random between 0 and 1500).\n\n", sum)
	return T
}

func (td *ThreeD) Start() {
	for {
		lock.Lock()
		listOfABS[td.Name]++
		times := td.Zero.Ping(td)
		for i, tron := range td.Trons {
			fmt.Printf("Delay %d milliseconds.\n", times[i])
			time.Sleep(time.Millisecond * time.Duration(times[i]))
			tron.Ping(td)
			listOfABS[td.Name]++
		}
		if listOfABS["a"] >= 4 && listOfABS["b"] >= 4 {
			cacheTimes = MakeTimes()
			listOfABS = map[string]int{}
		}
		lock.Unlock()
	}
}

func Start() {
	lock = &sync.Mutex{}
	zero := &Tron{}
	zero.Name = "zero"
	a := MakeThreeD(zero, "a")
	b := MakeThreeD(zero, "b")
	go a.Start()
	go b.Start()

	for {
		time.Sleep(time.Second * 1)
	}
}
