package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var nineP int
var nineN int
var ninePs []int = []int{}
var nineNs []int = []int{}
var sixThreePs []int = []int{}
var sixThreeNs []int = []int{}
var onePs []int = []int{}
var oneNs []int = []int{}

func main2() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("six separate infinities:")

	go ByNine(true)
	go ByNine(false)

	go BySixThree(true)
	go BySixThree(false)

	go ByOnes(true)
	go ByOnes(false)

	go Printer()

	for {
		time.Sleep(time.Second)
	}
}

func Printer() {
	for {
		time.Sleep(time.Second)
		buff := []string{}
		for _, val := range nineNs {
			buff = append([]string{fmt.Sprintf("%d", val)}, buff...)
		}
		buff = append(buff, "|")
		for _, val := range ninePs {
			buff = append(buff, fmt.Sprintf("%d", val))
		}
		fmt.Println(strings.Join(buff, " "))

		buff = []string{}
		for _, val := range sixThreeNs {
			buff = append([]string{fmt.Sprintf("%d", val)}, buff...)
		}
		buff = append(buff, "|")
		for _, val := range sixThreePs {
			buff = append(buff, fmt.Sprintf("%d", val))
		}
		fmt.Println(strings.Join(buff, " "))

		buff = []string{}
		for _, val := range oneNs {
			buff = append([]string{fmt.Sprintf("%d", val)}, buff...)
		}
		buff = append(buff, "|")
		for _, val := range onePs {
			buff = append(buff, fmt.Sprintf("%d", val))
		}
		fmt.Println(strings.Join(buff, " "))

	}
}

func ByNine(positive bool) {
	val := 9
	for {
		time.Sleep(time.Second)
		val = val * 2
		if positive {
			nineP = val
			ninePs = append(ninePs, nineP)
		} else {
			nineN = val * -1
			nineNs = append(nineNs, nineN)
		}
	}
}
func BySixThree(positive bool) {
	val := 3
	for {
		time.Sleep(time.Second)
		val = val * 2
		if positive {
			sixThreePs = append(sixThreePs, val)
		} else {
			sixThreeNs = append(sixThreeNs, val*-1)
		}

	}
}
func ByOnes(positive bool) {
	val := 1
	for {
		time.Sleep(time.Second)
		val = val * 2
		if positive {
			onePs = append(onePs, val)
		} else {
			oneNs = append(oneNs, val*-1)
		}
	}
}
