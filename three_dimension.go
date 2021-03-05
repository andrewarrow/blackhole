package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var lock *sync.Mutex

type Tron struct {
	Name string
}

func MakeTron(name string) *Tron {
	t := Tron{}
	t.Name = name
	return &t
}

func (t *Tron) Ping(from3D *ThreeD) {
	if t.Name == "zero" {
		fmt.Printf("The SHARED ZERO hit by %s\n", from3D.Name)
	} else {
		fmt.Printf("%s is at position %s\n", from3D.Name, t.Name)
	}
}

type ThreeD struct {
	Zero  *Tron
	Name  string
	Trons []*Tron
	Times []int
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

func (td *ThreeD) PickTimes() {
	td.Times = []int{}
	td.Times = append(td.Times, rand.Intn(500))
	td.Times = append(td.Times, rand.Intn(500))
	td.Times = append(td.Times, rand.Intn(500))
}

func (td *ThreeD) Start() {
	for {
		lock.Lock()
		td.Zero.Ping(td)
		lock.Unlock()
		td.PickTimes()
		for i, tron := range td.Trons {
			tron.Ping(td)
			time.Sleep(time.Millisecond * time.Duration(td.Times[i]))
		}
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
