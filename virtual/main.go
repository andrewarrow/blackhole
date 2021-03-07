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
	fmt.Println(fb.Base2Add(&fb))

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

	a = FancyBit{1}
	b = FancyBit{0}
	c = FancyBit{2}
	d = FancyBit{3}
	e := FancyBit{4}
	f := FancyBit{5}
	g := FancyBit{6}
	h := FancyBit{7}
	i := FancyBit{8}
	fb = FancyByte{}
	fb.List = append(fb.List, a)
	fb.List = append(fb.List, b)
	fb.List = append(fb.List, c)
	fb.List = append(fb.List, d)
	fb.List = append(fb.List, e)
	fb.List = append(fb.List, f)
	fb.List = append(fb.List, g)
	fb.List = append(fb.List, h)
	fb.List = append(fb.List, i)

	fb.Base9()
	fmt.Println(fb)
}
