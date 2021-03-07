package main

import (
	"fmt"
	"math"
)

type FancyBit struct {
	Val byte
}

type FancyByte struct {
	List []FancyBit
}

func (fb *FancyByte) Base2() {
	sum := 0
	for power, item := range fb.List {
		val := int(math.Exp2(float64(power)))
		sum += val * int(item.Val)
		fmt.Println(power, val, item, sum)
	}
}
func (fb *FancyByte) Base3() {
	sum := 0
	for power, item := range fb.List {
		val := int(math.Pow(3, float64(power)))
		sum += val * int(item.Val)
		fmt.Println(power, val, item, sum)
	}
}
func (fb *FancyByte) Base9() {
	sum := 0
	for power, item := range fb.List {
		val := int(math.Pow(9, float64(power)))
		sum += val * int(item.Val)
		fmt.Println(power, val, item, sum)
	}
}
