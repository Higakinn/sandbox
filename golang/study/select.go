// Go の _select_ を使うと、複数のチャネルで受信を待てる。
// 参考：　https://go.dev/play/p/Az1OY2XHN3a
// ゴルーチンとチャネル、そして select の組み合わせは Go の強力な機能である。
package main

import (
	"fmt"
	"time"
)

func main() {
	sum := make(chan int)
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go func() {
		sumValue := 0
		for _,v := range s {
			sumValue += v 
		} 
		sum <- sumValue
	}()

	go func() {
		for i := 0; i<=100; i++ {
			time.Sleep(1 * time.Millisecond)
			if i%2==0 {
				even<-i
			} else {
				odd<-i
			}
		}
		quit <- true
	}()

	func() {
		for {
			select {
			case v:= <- sum:
				fmt.Println("sum channel:", v)
			case v := <- even:
				fmt.Println("even channel:", v)
			case v := <- odd:
				fmt.Println("odd channel:", v)
			case <- quit:
				fmt.Println("quit")
				return
			}
		}
	}()
	fmt.Println("end!!")
}
