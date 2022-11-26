// go:build ignore
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	fmt.Printf("args count: %d\n", len(args))
	if len(args) < 3 {
		fmt.Println("Usage: go run study/byte.go 1 test")
		panic("argument not enough err")
	}
	b, _ := strconv.Atoi(args[1])
	bytes := int64(b)
	str := args[2]
	sr := strings.NewReader(str + "\n")
	lr := io.LimitReader(sr, bytes)
	cw, err := io.Copy(os.Stdout, lr);
	if err != nil {
		panic(err)
	}
	lr.
}
