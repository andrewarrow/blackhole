package main

type FourD struct {
	Name  string
	Zero  *Tron
	Trons []*Tron
}

func MakeFourD(zero *Tron, name string) *FourD {
	td := FourD{}
	td.Name = name
	td.Zero = zero
	td.Trons = []*Tron{}

	t1 := MakeTron("t1")
	t2 := MakeTron("t2")
	td.Trons = append(td.Trons, t1)
	td.Trons = append(td.Trons, t2)
	return &td
}

func (td *FourD) StartDraw() {
}
