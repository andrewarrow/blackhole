package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func Printer124() {
	list := []string{"1", "2", "4"}
	i := 0
	for {
		digit := list[i]
		fmt.Printf(digit)
		if rand.Intn(100) == 5 && digit == "2" {
			fmt.Printf("3")
		}
		time.Sleep(time.Millisecond * 100)
		i++
		if i == 3 {
			i = 0
		}
	}
}
func Printer875() {
	list := []string{"8", "7", "5"}
	i := 0
	for {
		digit := list[i]
		fmt.Printf(digit)
		if rand.Intn(100) == 5 && digit == "7" {
			fmt.Printf("6")
		}
		time.Sleep(time.Millisecond * 100)
		i++
		if i == 3 {
			i = 0
		}
	}
}
func Printer9() {
	for {
		fmt.Printf("9")
		time.Sleep(time.Second * 1000)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//go Printer124()
	//go Printer875()
	//go Printer9()

	bs, _ := ioutil.ReadFile("pattern.txt")
	for _, b := range bs {
		if b == 49 { // 1
			fmt.Println("<div style='display:inline-block;width:10px;height:10px;background-color:#FFAABB'> </div>")
		} else if b == 50 { // 2
			fmt.Println("<div style='display:inline-block;width:10px;height:10px;background-color:#00008B'> </div>")
		} else if b == 51 { // 3
			fmt.Println("<div style='display:inline-block;width:10px;height:10px;background-color:#FF7F50'> </div>")
		} else if b == 52 { // 4
			fmt.Println("<div style='display:inline-block;width:10px;height:10px;background-color:#FFE4C4'> </div>")
		} else if b == 53 { // 5
			fmt.Println("<div style='display:inline-block;width:10px;height:10px;background-color:#00FFFF'> </div>")
		} else if b == 54 { // 6
			fmt.Println("<div style='display:inline-block;width:10px;height:10px;background-color:#800080'> </div>")
		} else if b == 55 { // 7
			fmt.Println("<div style='display:inline-block;width:10px;height:10px;background-color:#008000'> </div>")
		} else if b == 56 { // 8
			fmt.Println("<div style='display:inline-block;width:10px;height:10px;background-color:#EE12DD'> </div>")
		} else if b == 59 { // 9
		}

	}

	//	time.Sleep(time.Second * 1)
	// 124
	// 124
	// 124
	// 1234
	// 124
	// 124
	// 124
	// three's happen, but rarely

	// 875
	// 875
	// 875
	// 8765
	// 875
	// 875
	// 875
	// six's happen, but rarely

	// 421
	// 421
	// 4231
	// 421
	// 421

	// 578
	// 578
	// 5678
	// 578
	// 578

}

/*
func main3() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("six separate infinities:")

	go ByNine(true)
	go ByNine(false)

	go BySixThree(true)
	go BySixThree(false)

	go ByOnes(true)
	go ByOnes(false)

	go Printer()

	for {
		time.Sleep(time.Second)
	}
}

func Printer() {
	for {
		time.Sleep(time.Second)
		buff := []string{}
		for _, val := range nineNs {
			buff = append([]string{fmt.Sprintf("%d", val)}, buff...)
		}
		buff = append(buff, "|")
		for _, val := range ninePs {
			buff = append(buff, fmt.Sprintf("%d", val))
		}
		fmt.Println(strings.Join(buff, " "))

		buff = []string{}
		for _, val := range sixThreeNs {
			buff = append([]string{fmt.Sprintf("%d", val)}, buff...)
		}
		buff = append(buff, "|")
		for _, val := range sixThreePs {
			buff = append(buff, fmt.Sprintf("%d", val))
		}
		fmt.Println(strings.Join(buff, " "))

		buff = []string{}
		for _, val := range oneNs {
			buff = append([]string{fmt.Sprintf("%d", val)}, buff...)
		}
		buff = append(buff, "|")
		for _, val := range onePs {
			buff = append(buff, fmt.Sprintf("%d", val))
		}
		fmt.Println(strings.Join(buff, " "))

	}
}

func ByNine(positive bool) {
	val := 9
	for {
		time.Sleep(time.Second * 365)
		val = val * 2
		if positive {
			nineP = val
			ninePs = append(ninePs, nineP)
		} else {
			nineN = val * -1
			nineNs = append(nineNs, nineN)
		}
	}
}
func BySixThree(positive bool) {
	val := 3
	for {
		time.Sleep(time.Second * 1)
		val = val * 2
		if positive {
			sixThreePs = append(sixThreePs, val)
		} else {
			sixThreeNs = append(sixThreeNs, val*-1)
		}

	}
}
func ByOnes(positive bool) {
	val := 1
	for {
		time.Sleep(time.Millisecond * 1)
		val = val * 2
		if positive {
			onePs = append(onePs, val)
		} else {
			oneNs = append(oneNs, val*-1)
		}
	}
}
var nineP int
var nineN int
var ninePs []int = []int{}
var nineNs []int = []int{}
var sixThreePs []int = []int{}
var sixThreeNs []int = []int{}
var onePs []int = []int{}
var oneNs []int = []int{}
*/
