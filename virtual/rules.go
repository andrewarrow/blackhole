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

	// 3936936393363
	// 9366639393936
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
