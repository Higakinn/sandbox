package main

import (
	"fmt"
	"math"
)

func primeNumber() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		for i := 3; i < 100000; i+=2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j < l + 1; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			} 
			if !found {
				result <- i
			}
		}
		close(result)
	}()
	return result
}

func evenNumber() chan int {
	r := make(chan int)
	go func() {
		for i := 0; i <= 100; i += 1 {
			if i % 2 == 0 {
				r <- i
			}
		}
		close(r)
	}()
	return r
}

func main() {
	en := evenNumber()
	for n := range en {
		fmt.Println(n)
	}
	// pn := primeNumber()
	// for n := range pn {
	// 	fmt.Println(n)
	// } 
}