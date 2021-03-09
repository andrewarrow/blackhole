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
	fmt.Println("9 & -9")
	fmt.Println("13 paths(s): -1,-2,-4,-5,-7,-8, 0, 1, 2, 4, 5, 7, 8")
	fmt.Println("-> -1")
	fmt.Println("18 from -1")
	fmt.Println(" x paths(s): -1,19,19,-1,-1,0,0,-1")
	fmt.Println("-> -2")
	fmt.Println("18 from -2")
	fmt.Println(" x paths(s): -1,-1,20,-1,20,-1,20,-1,-1")
	fmt.Println(" x paths(s): -2,20,20,-2,0,0,-2")
	fmt.Println("-> -4")
	fmt.Println("18 from -4")
	fmt.Println(" x paths(s): -1,-1,-1,-1,22,22,-1,-1,-1,-1,0,0,-1,-1,-1,-1")
	fmt.Println(" x paths(s): -2,-2,22,22,-2,-2,0,0,-2,-2")
	fmt.Println(" x paths(s): -4,22,22,-4,0,0,-4")
	fmt.Println("18 from -5")
	fmt.Println(" x paths(s): -1,-1,-1,-1,-1,23,23,-1,-1,-1,-1,-1,0,0,-1,-1,-1,-1,-1")
	fmt.Println(" x paths(s): -2,-2,-1,23,23,-2,-2,-1,0,0,-2,-2,-1")
	fmt.Println(" x paths(s): -4,-1,23,23,-4,-1,0,0,-4,-1")
	fmt.Println(" x paths(s): -5,23,23,-5,0,0,-5")
	fmt.Println("36 from -1,19")
	fmt.Println(" x paths(s): -1,19,18")
	fmt.Println(" x paths(s): -1,19,12,3,3")
	fmt.Println(" x paths(s): -1,19,12,3,3")
	fmt.Println(" x paths(s): -1,19,12,6")
	fmt.Println(" x paths(s): -1,19,9,9")
	fmt.Println(" x paths(s): -1,19,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1")
	//     1+17,2+16,3+15,4+14,5+13,6+12,7+11
	//     8+10,9+9,10+8,11+7,12+6,13+5,14+4,
	//     15+3,16+2,17+1,6+6+6,3+6+6+3,6+6+3+3,
	//     3+9+3+3,9+3+3+3,3+3+3+9
	//     24-6,23-6+1
}
func main3() {
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
		if usSeed == 1280000000000000 {
			usSeed = 1
		}
		if sign {
			fmt.Println(usSeed, loops)
		} else {
			fmt.Println("-", usSeed, loops)
		}
		usSeed = usSeed * 2
		// 1,2,4,8,16,32,64
	}
}
func maleFemale() {
	loops := 0
	for sign := range uniCrank {
		if mfSeed == 480000000000 {
			mfSeed = 3
		}
		loops++
		if sign {
			fmt.Println("                     ", mfSeed, loops)
		} else {
			fmt.Println("                    -", mfSeed, loops)
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
			fmt.Println("                                 ", int64(num)*uniSeed, loops)
			uniCrank <- num > 0
		}
		loops++
		uniSeed = uniSeed * 2
		time.Sleep(time.Millisecond * 100)
	}
}
