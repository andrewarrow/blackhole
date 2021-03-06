package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var lockMF sync.Mutex = sync.Mutex{}
var fourCrank chan bool = make(chan bool, 1024)
var uniCrank chan bool = make(chan bool, 1024)
var uniSeed = 9
var mfSeed = 3
var usSeed = 1

func main() {
	rand.Seed(time.Now().UnixNano())
	go three()
	go maleFemale()
	go theUni()
	for {
		time.Sleep(time.Second)
	}
}
func three() {
	loops := 0
	for _ = range fourCrank {
		list := []int{1, 2, 4, 0, 8, 7, 5, 0, -1, -2, -4, 0, -8, -7, -5, 0}
		for _, num := range list {
			if num == 8 || num == 7 || num == 5 {
				fmt.Println(usSeed, loops, "LEFT", num)
			} else if num == 0 {
				fmt.Println(usSeed, loops, "ZERO", num)
			} else {
				fmt.Println(usSeed, loops, "RIGHT", num)
			}
		}
		loops++
		usSeed = usSeed * 2
	}
}
func maleFemale() {
	loops := 0
	for _ = range uniCrank {
		list := []int{3, 6, 9, -3, -6}
		for _, num := range list {
			if num == 9 {
				lockMF.Lock()
				lockMF.Unlock()
				continue
			}
			fmt.Println("                  ", mfSeed, loops, num)
		}
		loops++
		mfSeed = mfSeed * 2
		fourCrank <- true
	}
}
func theUni() {
	loops := 0
	for {
		list := []int{9}
		for _, num := range list {
			lockMF.Lock()
			fmt.Println("                                 ", uniSeed, loops, num)
			lockMF.Unlock()
		}
		loops++
		uniSeed = uniSeed * 2
		uniCrank <- true
		time.Sleep(time.Second)
	}
}
