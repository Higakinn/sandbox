package main

import (
	"fmt"
)

type Talker interface {
	Talk()
}

type Greeter struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func hoge() {
	var talker Talker
	talker = &Greeter{"woxoxo"}
	talker.Talk()
}
