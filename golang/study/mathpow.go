/*
* べき乗を計算するためのサンプルプログラム
 */
// go:build
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var args []float64
	// コマンドライン引数(str)をfloat64に変換する
	for _, arg := range os.Args[1:] {
		if n, err := strconv.ParseFloat(arg, 64); err == nil {
			args = append(args, n)
		}
	}
	if len(args) <= 1 || len(args) >= 4 {
		fmt.Println("Usage: go run study/mathpow.go [radix(基数)] [index(指数)]")
		panic("arg error")
	}
	//結果を標準出力で表示
	fmt.Println(math.Pow(args[0], args[1]))
}
