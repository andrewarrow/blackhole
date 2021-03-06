package main

import (
	"fmt"
	"time"
)

func main() {
	go three()
	go four()
	for {
		time.Sleep(time.Second)
	}
}
func three() {
	loops := 0
	evenOdd := false
	for {
		list := []byte{1, 2, 4, 0, 8, 7, 5, 0}
		for _, num := range list {
			if evenOdd || num == 0 {
				fmt.Println(loops, num)
			} else {
				fmt.Println(loops, "-", num)
			}
		}
		loops++
		evenOdd = !evenOdd
	}
}
func four() {
	loops := 0
	evenOdd := true
	for {
		list := []int{3, 9, 6, -3, -9, -6}
		for _, num := range list {
			fmt.Println("                  ", loops, num)
		}
		loops++
		evenOdd = !evenOdd
	}
}
