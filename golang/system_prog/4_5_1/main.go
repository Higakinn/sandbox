package main

import (
	"fmt"
	"time"
)

var c chan int

func handle(int) {}

func main() {
	select {
	case m := <-c:
		fmt.Println(m)
		handle(m)
	case <-time.After(10 * time.Second):
		fmt.Println("timed out")
	}
}