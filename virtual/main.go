package main

import (
	"fmt"
)

func main() {
	a := FancyBit{1}
	b := FancyBit{0}
	c := FancyBit{1}
	fb := FancyByte{}
	fb.List = append(fb.List, a)
	fb.List = append(fb.List, b)
	fb.List = append(fb.List, c)

	fb.Base2()
	fmt.Println(fb)

	a = FancyBit{2}
	b = FancyBit{0}
	c = FancyBit{2}
	d := FancyBit{1}
	fb = FancyByte{}
	fb.List = append(fb.List, a)
	fb.List = append(fb.List, b)
	fb.List = append(fb.List, c)
	fb.List = append(fb.List, d)

	fb.Base3()
	fmt.Println(fb)
}
