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
		if listOfABS4["c"] >= 3 && listOfABS4["d"] >= 3 {
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
		if from4D.Name == "c" {
			DrawWithParams(DrawParams{4, "c", 0})
		} else {
			DrawWithParams(DrawParams{4, "c", 0})
		}
	} else {
		if from4D.Name == "d" {
			DrawWithParams(DrawParams{4, "d", listOfABS4[from4D.Name] + 1})
		} else {
			DrawWithParams(DrawParams{4, "d", listOfABS4[from4D.Name] + 1})
		}
	}
	return list
}
