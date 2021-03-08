package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var sharedZero sync.Mutex = sync.Mutex{}
var lock sync.Mutex = sync.Mutex{}
var fourChan chan bool = make(chan bool, 1024)

var uniCrank chan bool = make(chan bool, 1024)
var uniSeed = int64(9)
var mfSeed = int64(3)
var usSeed = int64(1)

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
	for sign := range fourChan {
		loops++
		if usSeed == 128 {
			usSeed = 1
		}
		if sign {
			fmt.Println(loops, usSeed)
		} else {
			fmt.Println(loops, "-", usSeed)
		}
		usSeed = usSeed * 2
		// 1,2,4,8,16,32,64
	}
}
func maleFemale() {
	loops := 0
	for sign := range uniCrank {
		if mfSeed == 48 {
			mfSeed = 3
		}
		loops++
		if sign {
			fmt.Println("                     ", loops, mfSeed)
		} else {
			fmt.Println("                    -", loops, mfSeed)
		}
		mfSeed = mfSeed * 2
		fourChan <- sign
	}
}
func theUni() {
	loops := 0
	for {
		list := []int{1, -1}
		for _, num := range list {
			fmt.Println("                                 ", int64(num)*uniSeed, loops, num)
			uniCrank <- num > 0
		}
		loops++
		uniSeed = uniSeed * 2
		time.Sleep(time.Millisecond * 100)
	}
}
