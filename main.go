package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Robert Sepehr's Virtual Track Computer")
	fmt.Println("")
	fmt.Println("Since all roads lead to rome there is no escpaing...")
	fmt.Println("")
	fmt.Println("The Blackhole.")
	fmt.Println("")
	fmt.Println("")
	talk := `Eventually the earth will be pulled crashing it flat into
the center of the milky way galaxy solar system.

The earth mountains will be pulled crushing flat, 
the seas will boil off, the earth will be mixed together
with other planets and stars spun around faster and faster until...
we all shoot out!

The other side of the blackhole is a hot gas thing of nebular matter.
It will go billions of lightyears in the form of a way-ho which is
a pulsar.

Every black hole has a way-ho on the other side.

It will eventually emalgomate, new planets new stars, why does this happen?
It's a process of renewal.

A blackhole is a compost pile.
	`
	fmt.Println(talk)
	rand.Seed(time.Now().UnixNano())
	StartWithDrawing()

}
