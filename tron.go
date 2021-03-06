package main

import (
	"fmt"
	"sync"
)

var listOfABS3 map[string]int = map[string]int{}
var listOfABS4 map[string]int = map[string]int{}
var cacheTimes []int = MakeTimes(false)
var lock sync.Mutex = sync.Mutex{}
var zeroFor3 *Tron
var zeroFor4 *Tron
var totalSums int
var totalCount int

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
		list = cacheTimes
	} else {
		fmt.Printf("%s is at position %s\n", from3D.Name, t.Name)
	}
	return list
}
