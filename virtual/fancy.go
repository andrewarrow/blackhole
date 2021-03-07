package main

import (
	"fmt"
	"math"
)

/*                                             off and off = off
0 1 {1} 1                  1|1=1   1&1=1       on  or off = on
1 2 {0} 1                  1|0=1   1&0=0       on and off = off
2 4 {1} 5                  0|0=0   0&0=0       on  or  on = on
{[{1} {0} {1}]}
0 1 {2} 2                  1|1|1=2   2|2|2=2   on or off or middle = on
1 3 {0} 2                  1|2|0=2             on and off and middle = middle
2 9 {2} 20                 1|0|1=2             on and off or middle = off
3 27 {1} 47
{[{2} {0} {2} {1}]}
0 1 {1} 1
1 9 {0} 1
2 81 {2} 163
3 729 {3} 2350
4 6561 {4} 28594
5 59049 {5} 323839
6 531441 {6} 3512485
7 4782969 {7} 36993268
8 43046721 {8} 381367036
{[{1} {0} {2} {3} {4} {5} {6} {7} {8}]}

*/
type FancyBit struct {
	Val byte
}

type FancyByte struct {
	List []FancyBit
}

func MakeBase3FancyByte(s string) FancyByte {
	result := FancyByte{}

	a9 := FancyBit{9}
	a3 := FancyBit{3}
	a6 := FancyBit{6}
	for a := range s {
		if s[a] == 57 {
			result.List = append(result.List, a9)
		} else if s[a] == 51 {
			result.List = append(result.List, a3)
		} else if s[a] == 54 {
			result.List = append(result.List, a6)
		}
	}

	return result
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
