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
				fmt.Println(loops, "LEFT", num)
			} else if num == 0 {
				fmt.Println(loops, "ZERO", num)
			} else {
				fmt.Println(loops, "RIGHT", num)
			}
		}
		loops++
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
			fmt.Println("                  ", loops, num)
		}
		loops++
		fourCrank <- true
	}
}
func theUni() {
	loops := 0
	for {
		list := []int{9}
		for _, num := range list {
			lockMF.Lock()
			fmt.Println("999999999999999999              ", loops, num)
			lockMF.Unlock()
		}
		loops++
		uniCrank <- true
		time.Sleep(time.Second)
	}
}
