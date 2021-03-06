package main

import (
	"fmt"
	"math/rand"
	"time"
)

var fourCrank chan bool = make(chan bool, 1024)

func main() {
	rand.Seed(time.Now().UnixNano())
	go three()
	go four()
	for {
		time.Sleep(time.Second)
	}
}
func three() {
	loops := 0
	for _ = range fourCrank {
		list := []int{1, 2, 4, 0, -8, -7, -5, 0, -1, -2, -4, 0, 8, 7, 5, 0}
		for _, num := range list {
			if rand.Intn(990000) == 19 {
				fmt.Println(loops, num)
			}
		}
		loops++
	}
}
func four() {
	loops := 0
	for {
		list := []int{3, 9, 6, -3, -9, -6}
		for _, num := range list {
			if rand.Intn(990000) == 19 {
				fmt.Println("                  ", loops, num)
			}
		}
		loops++
		fourCrank <- true
	}
}
