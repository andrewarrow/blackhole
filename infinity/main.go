package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("three separate infinities:")

	go ByNine(true)
	go ByNine(false)

	go BySixThree(true)
	go BySixThree(false)

	go ByOnes(true)
	go ByOnes(false)

	for {
		time.Sleep(time.Second)
	}
}

func ByNine(positive bool) {
	val := 9
	for {
		time.Sleep(time.Second)
		val = val * 2
		if positive {
			fmt.Println("nine", " ", val)
		} else {
			fmt.Println("nine", "-", val)
		}
	}
}
func BySixThree(positive bool) {
	val := 3
	for {
		time.Sleep(time.Second)
		val = val * 2
		if positive {
			fmt.Println("sixThree", " ", val)
		} else {
			fmt.Println("sixThree", "-", val)
		}

	}
}
func ByOnes(positive bool) {
	val := 1
	for {
		time.Sleep(time.Second)
		val = val * 2
		if positive {
			fmt.Println("ones", " ", val)
		} else {
			fmt.Println("ones", "-", val)
		}
	}
}
