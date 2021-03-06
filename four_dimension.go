package main

import "time"

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

func (fd *FourD) StartDraw() {
	for {
		lock.Lock()
		listOfABS4[fd.Name]++
		time.Sleep(time.Millisecond * 100)
		times := fd.Zero.PingDraw4(fd)
		for i, tron := range fd.Trons {
			time.Sleep(time.Millisecond * time.Duration(times[i]))
			tron.PingDraw4(fd)
			listOfABS4[fd.Name]++
		}
		if listOfABS4["a"] >= 3 && listOfABS4["b"] >= 3 {
			cacheTimes = MakeTimes(false)
			listOfABS4 = map[string]int{}
			//completedRevs++
		}
		lock.Unlock()
	}
}
func (t *Tron) PingDraw4(from4D *FourD) []int {
	list := []int{}
	if t.Name == "zero" {
		list = cacheTimes
		if from4D.Name == "a" {
			DrawWithParams(DrawParams{"a", 0})
		} else {
			DrawWithParams(DrawParams{"b", 0})
		}
	} else {
		if from4D.Name == "a" {
			DrawWithParams(DrawParams{"a", listOfABS4[from4D.Name] + 1})
		} else {
			DrawWithParams(DrawParams{"b", listOfABS4[from4D.Name] + 1})
		}
	}
	return list
}
