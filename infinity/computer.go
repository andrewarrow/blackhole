package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Computer struct {
	TablesP []int
	TablesN []int
}

func (c *Computer) Run() {
	i := 0
	direction := false
	for {
		time.Sleep(time.Millisecond * 400)
		from := "     SUN ---======********* "
		if c.TablesP[0] == 6 {
			from = "     DAY ---====== "
		} else if c.TablesP[0] == 2 {
			from = "ROTATION "
		}
		fmt.Println(from, c.TablesP[i], c.TablesN[i])
		if direction {
			i--
		} else {
			i++
		}
		if i == 0 {
			direction = false
			continue
		}
		if i >= len(c.TablesP) {
			i--
			direction = true
			continue
		}
		if i >= len(c.TablesN) {
			i--
			direction = true
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("six separate infinities:")
	c1 := MakeLevel1Computer()
	c3 := MakeLevel3Computer()
	c9 := MakeLevel9Computer()

	go c1.Run()
	go c3.Run()
	go c9.Run()
	for {
		time.Sleep(time.Second * 1)
	}
}

func MakeLevel1Computer() *Computer {
	c := Computer{}
	c.TablesP = []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304, 8388608, 16777216, 33554432, 67108864, 134217728, 268435456, 536870912, 1073741824, 2147483648, 4294967296, 8589934592, 17179869184, 34359738368, 68719476736, 137438953472, 274877906944, 549755813888, 1099511627776}
	c.TablesN = []int{-1099511627776, -549755813888, -274877906944, -137438953472, -68719476736, -34359738368, -17179869184, -8589934592, -4294967296, -2147483648, -1073741824, -536870912, -268435456, -134217728, -67108864, -33554432, -16777216, -8388608, -4194304, -2097152, -1048576, -524288, -262144, -131072, -65536, -32768, -16384, -8192, -4096, -2048, -1024, -512, -256, -128, -64, -32, -16, -8, -4, -2}
	return &c
}

func MakeLevel3Computer() *Computer {
	c := Computer{}
	c.TablesP = []int{6, 12, 24, 48, 96, 192, 384, 768, 1536, 3072, 6144, 12288, 24576, 49152, 98304, 196608, 393216, 786432, 1572864, 3145728, 6291456, 12582912, 25165824, 50331648, 100663296, 201326592, 402653184, 805306368, 1610612736, 3221225472, 6442450944, 12884901888, 25769803776, 51539607552, 103079215104, 206158430208, 412316860416, 824633720832, 1649267441664, 3298534883328}
	c.TablesN = []int{-3298534883328, -1649267441664, -824633720832, -412316860416, -206158430208, -103079215104, -51539607552, -25769803776, -12884901888, -6442450944, -3221225472, -1610612736, -805306368, -402653184, -201326592, -100663296, -50331648, -25165824, -12582912, -6291456, -3145728, -1572864, -786432, -393216, -196608, -98304, -49152, -24576, -12288, -6144, -3072, -1536, -768, -384, -192, -96, -48, -24, -12, -6}
	return &c
}
func MakeLevel9Computer() *Computer {
	c := Computer{}
	c.TablesP = []int{18, 36, 72, 144, 288, 576, 1152, 2304, 4608, 9216, 18432, 36864, 73728, 147456, 294912, 589824, 1179648, 2359296, 4718592, 9437184, 18874368, 37748736, 75497472, 150994944, 301989888, 603979776, 1207959552, 2415919104, 4831838208, 9663676416, 19327352832, 38654705664, 77309411328, 154618822656, 309237645312, 618475290624, 1236950581248, 2473901162496, 4947802324992, 9895604649984}
	c.TablesN = []int{-4947802324992, -2473901162496, -1236950581248, -618475290624, -309237645312, -154618822656, -77309411328, -38654705664, -19327352832, -9663676416, -4831838208, -2415919104, -1207959552, -603979776, -301989888, -150994944, -75497472, -37748736, -18874368, -9437184, -4718592, -2359296, -1179648, -589824, -294912, -147456, -73728, -36864, -18432, -9216, -4608, -2304, -1152, -576, -288, -144, -72, -36, -18}
	return &c
}
