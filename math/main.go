package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var lock sync.Mutex = sync.Mutex{}
var fourCrank chan bool = make(chan bool, 1024)

func main() {
	rand.Seed(time.Now().UnixNano())
	go three1()
	go three2()
	go maleFemale()
	go theUni()
	for {
		time.Sleep(time.Second)
	}
}
func three2() {
	loops := 0
	for _ = range fourCrank {
		list := []int{1, 2, 4, 0, -4, -2, -1, 0}
		for _, num := range list {
			if rand.Intn(990000) == 19 {
				fmt.Println(loops, "LEFT", num)
			}
		}
		loops++
	}
}
func three1() {
	loops := 0
	for _ = range fourCrank {
		list := []int{8, 7, 5, 0, -5, -7, -8, 0}
		for _, num := range list {
			if rand.Intn(990000) == 19 {
				fmt.Println(loops, "RIGHT", num)
			}
		}
		loops++
	}
}
func maleFemale() {
	loops := 0
	for _ = range fourCrank {
		list := []int{3, 6, 9, -3, -6}
		for _, num := range list {
			if num == 9 {
				lock.Lock()
				lock.Unlock()
				continue
			}
			if rand.Intn(90000) == 19 {
				fmt.Println("                  ", loops, num)
			}
		}
		loops++
	}
}
func theUni() {
	loops := 0
	for {
		list := []int{9}
		for _, num := range list {
			lock.Lock()
			if rand.Intn(990000) == 19 {
				fmt.Println("999999999999999999              ", loops, num)
			}
			lock.Unlock()
		}
		loops++
		fourCrank <- true
	}
}
