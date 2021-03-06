package main

import "fmt"

/*
6561, 13122
  81,   162, 324, 648, 1296, 2592, 5148
  9,     18,  36,  72,  144,  288,  576
  3,      6,  12,  24,   48,   96,  192
  1,      2,   4,   8,   16,   32,   64


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
//
// we see the shadow because this event HAPPENED. It already happened in 9D first,
// then 4D, so when we are left with N possible 4D paths we know one selection WAS
// ALREADY made. We just don't know which. And when we observe, one is revealed. But
// act of observation isn't making the selection. The selection already happened
// and we have the shadow to prove it.
//
// time flow in level           9 ---------->
//      time flow in level      3 --------->
//           time flow in level 1 -------->
//
// 9 is in the lead, always.
// In our 1,2,4,8,7,5 world of level 1, every cycle we pass thru 0 means
// nothing is added from our "bad" dimension. We can experience something direct
// from level 3 or every so often, not only is our cycle at 0, but level 3's
// cycle is exactly 50% between 3 and 6. Talk about stars lining up.
//
// When we experience pure 9, it's OMG amazing.
// Even just pure level 3 is OMG amazing compared to level 1.
// Level 1 is the lowest you can get.
//
// 1,2,4=male
// 5,7,8=female
//

func main2() {
	list := TryAndMake(1052, []string{})
	fmt.Println(list)
}
