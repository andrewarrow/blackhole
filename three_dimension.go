package main

import (
	"fmt"
	"time"
)

type Tron struct {
}

func (t *Tron) Ping(ab string) {
	fmt.Printf("Ping from %s\n", ab)
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
	return &td
}

func (td *ThreeD) Start() {
	for {
		td.Zero.Ping(td.Name)
		time.Sleep(time.Second * 1)
	}
}

func Start() {
	zero := &Tron{}
	a := MakeThreeD(zero, "a")
	b := MakeThreeD(zero, "b")
	go a.Start()
	go b.Start()

	for {
		time.Sleep(time.Second * 1)
	}
}
