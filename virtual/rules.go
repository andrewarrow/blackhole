package main

import "fmt"

func (fb *FancyByte) Base2Add(target *FancyByte) *FancyByte {
	result := FancyByte{}

	// 1001010101001
	// 0110010010101
	// =
	// 1111100111110
	a1 := FancyBit{1}
	a0 := FancyBit{0}
	carry := false
	for i, n := range fb.List {
		if carry {
			result.List = append([]FancyBit{a1}, result.List...)
			carry = false
			continue
		}
		other := target.List[i].Val
		if other == 1 && n.Val == 1 {
			carry = true
			result.List = append([]FancyBit{a0}, result.List...)
		} else if other == 1 && n.Val == 0 {
			result.List = append([]FancyBit{a1}, result.List...)
		} else if other == 0 && n.Val == 1 {
			result.List = append([]FancyBit{a1}, result.List...)
		} else if other == 0 && n.Val == 0 {
			result.List = append([]FancyBit{a0}, result.List...)
		}
		fmt.Println(other, n)
	}
	if carry {
		result.List = append([]FancyBit{a1}, result.List...)
	}

	return &result
}

func (fb *FancyByte) Base3Add(target *FancyByte) *FancyByte {
	result := FancyByte{}

	// 63
	// 36
	// 9

	// 3936936393363
	// 9366639393936
	// =
	// 9966969696966
	a9 := FancyBit{9}
	a3 := FancyBit{3}
	a6 := FancyBit{6}
	carry := ""
	for i, n := range fb.List {
		other := target.List[i].Val
		if carry == "" {
			if other == 9 && n.Val == 9 {
				carry = "9"
				result.List = append([]FancyBit{a9}, result.List...)
			} else if other == 3 && n.Val == 3 {
				result.List = append([]FancyBit{a6}, result.List...)
			} else if other == 6 && n.Val == 6 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "3"
			} else if other == 6 && n.Val == 3 {
				result.List = append([]FancyBit{a9}, result.List...)
			} else if other == 3 && n.Val == 6 {
				result.List = append([]FancyBit{a9}, result.List...)
			} else if other == 9 && n.Val == 6 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "6"
			} else if other == 6 && n.Val == 9 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "6"
			} else if other == 3 && n.Val == 9 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "3"
			} else if other == 9 && n.Val == 3 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "3"
			}
		} else if carry == "3" {
			if other == 9 && n.Val == 9 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "3"
			} else if other == 3 && n.Val == 3 {
				result.List = append([]FancyBit{a9}, result.List...)
			} else if other == 6 && n.Val == 6 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "6"
			} else if other == 6 && n.Val == 3 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "3"
			} else if other == 3 && n.Val == 6 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "3"
			} else if other == 9 && n.Val == 6 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "9"
			} else if other == 6 && n.Val == 9 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "9"
			} else if other == 3 && n.Val == 9 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "6"
			} else if other == 9 && n.Val == 3 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "6"
			}
		} else if carry == "6" {
			if other == 9 && n.Val == 9 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "6"
			} else if other == 3 && n.Val == 3 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "3"
			} else if other == 6 && n.Val == 6 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "9"
			} else if other == 6 && n.Val == 3 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "6"
			} else if other == 3 && n.Val == 6 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "6"
			} else if other == 9 && n.Val == 6 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "3"
			} else if other == 6 && n.Val == 9 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "3"
			} else if other == 3 && n.Val == 9 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "9"
			} else if other == 9 && n.Val == 3 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "9"
			}
		} else if carry == "9" {
			if other == 9 && n.Val == 9 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "9"
			} else if other == 3 && n.Val == 3 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "6"
			} else if other == 6 && n.Val == 6 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "3"
			} else if other == 6 && n.Val == 3 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "9"
			} else if other == 3 && n.Val == 6 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "9"
			} else if other == 9 && n.Val == 6 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "6"
			} else if other == 6 && n.Val == 9 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "6"
			} else if other == 3 && n.Val == 9 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "3"
			} else if other == 9 && n.Val == 3 {
				result.List = append([]FancyBit{a9}, result.List...)
				carry = "3"
			}
		}
	}
	if carry != "" {
		if carry == "9" {
			result.List = append([]FancyBit{a9}, result.List...)
		} else if carry == "3" {
			result.List = append([]FancyBit{a3}, result.List...)
		} else if carry == "6" {
			result.List = append([]FancyBit{a6}, result.List...)
		}
	}

	return &result
}
