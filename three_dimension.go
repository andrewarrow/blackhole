package main

import (
	"fmt"
	"strings"
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

func (t *Tron) Ping(ab string) {
	if strings.HasSuffix(ab, "b") {
		fmt.Printf("                          Ping from %s, i'm %s\n", ab, t.Name)
	} else {
		fmt.Printf("Ping from %s, i'm %s\n", ab, t.Name)
	}
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

func (td *ThreeD) Start() {
	for {
		lock.Lock()
		td.Zero.Ping(td.Name)
		lock.Unlock()
		for i, tron := range td.Trons {
			tron.Ping(fmt.Sprintf("%d %s", i, td.Name))
		}
		time.Sleep(time.Second * 1)
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
