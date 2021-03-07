package main

import "fmt"

func (fb *FancyByte) Base2Add(target *FancyByte) *FancyByte {
	result := FancyByte{}

	// 1001010101001
	// 0110010010101
	// =
	// 1111100111110
	carry := false
	for i, n := range fb.List {
		if carry {
			result.List = append([]FancyBit{FancyBit{1}}, result.List...)
			carry = false
			continue
		}
		other := target.List[i].Val
		if other == 1 && n.Val == 1 {
			carry = true
			result.List = append([]FancyBit{FancyBit{0}}, result.List...)
		} else if other == 1 && n.Val == 0 {
			result.List = append([]FancyBit{FancyBit{1}}, result.List...)
		} else if other == 0 && n.Val == 1 {
			result.List = append([]FancyBit{FancyBit{1}}, result.List...)
		} else if other == 0 && n.Val == 0 {
			result.List = append([]FancyBit{FancyBit{0}}, result.List...)
		}
		fmt.Println(other, n)
	}
	if carry {
		result.List = append([]FancyBit{FancyBit{1}}, result.List...)
	}

	return &result
}
