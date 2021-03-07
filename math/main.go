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

/*

9, 18, 36, 72, 144, 288, 576
3,  6, 12, 24,  48,  96, 192
1,  2,  4,  8,  16,  32,  64


try and make 295, 288+6+1
try and make 201, 144+36=180+18=198+3
try and make 201, 144+48=192+9



*/
func TryAndMake(target int, list []string) []string {
	fmt.Println("tam", target, list)
	brokeAt := 0
	for i, n := range level9 {
		if n >= target {
			brokeAt = i
			break
		}
	}
	fmt.Println("ba", target, brokeAt)
	if brokeAt == 0 {
		nextBreak := 0
		for i, n := range level3 {
			if n >= target {
				nextBreak = i
				break
			}
		}
		if nextBreak == 0 {
			finalBreak := 0
			for i, n := range level1 {
				if n >= target {
					finalBreak = i
					break
				}
			}
			if finalBreak == 0 {
				return append(list, "1")
			}
			val := level1[finalBreak-1]
			// fyi when no return here, diff path
			return TryAndMake(target-val, append(list, fmt.Sprintf("%d", val)))
		}
		if nextBreak == 0 {
			return append(list, "3")
		}
		val := level3[nextBreak-1]
		return TryAndMake(target-val, append(list, fmt.Sprintf("%d", val)))
	}
	val := level9[brokeAt-1]
	return TryAndMake(target-val, append(list, fmt.Sprintf("%d", val)))
}

var level9 []int = []int{9, 18, 36, 72, 144, 288, 576, 1152}
var level3 []int = []int{3, 6, 12, 24, 48, 96, 192}
var level1 []int = []int{1, 2, 4, 8, 16, 32, 64}

// what assumptions can we make about the level3 and level9?
// We can see shadow of 4D object but not object itself
// That shadow from 4D, and the shadow of 9D in 4D, can be computed.
// So you can limit the possible routes you took in level3 and level9 but
// you can never know 100% sure. There are always multiple routes the
// shadow doesn't eliminate false route, only some, many perhaps, but not all.
//
// so, for a very large number we can say where it did NOT come from the 4D
// and we can say where it did NOT come from the 9D.

func main2() {
	list := TryAndMake(1052, []string{})
	fmt.Println(list)
}

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
