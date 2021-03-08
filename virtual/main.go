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
	//fmt.Println(fb.Base2Add(&fb))
	fb = MakeBase3FancyByte("3339")
	fb.Base3()

	// 12
	// 9+(3)
	// 9+3
	//
	//fb1 := MakeBase3FancyByte("3936936393363")
	fb1 := MakeBase3FancyByte("369")
	fb1.Base3()
	fmt.Println(fb1)
	//fb2 := MakeBase3FancyByte("9366639393936")
	fb2 := MakeBase3FancyByte("639")
	fb2.Base3()
	fmt.Println(fb2)
	sum := fb1.Base3Add(&fb2)
	sum.Base3()

	/*
		a = FancyBit{1}
		b = FancyBit{0}
		c = FancyBit{2}
		d := FancyBit{3}
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
		//fmt.Println(fb)*/
}
