package main

import (
	"fmt"
	"sync"
)

var listOfABS3 map[string]int = map[string]int{}
var listOfABS4 map[string]int = map[string]int{}
var cacheTimes3 []int = MakeTimes3(false)
var cacheTimes4 []int = MakeTimes4(false)
var lock sync.Mutex = sync.Mutex{}
var zeroFor3 *Tron
var zeroFor4 *Tron
var totalSums3 int
var totalCount3 int
var totalSums4 int
var totalCount4 int
var completedRevs3 int
var completedRevs4 int
var lastThing *TemplateThing

func init() {
	zeroFor3 = &Tron{}
	zeroFor3.Name = "zero"
	zeroFor4 = &Tron{}
	zeroFor4.Name = "zero"
}

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
		list = cacheTimes3
	} else {
		fmt.Printf("%s is at position %s\n", from3D.Name, t.Name)
	}
	return list
}
