package main

import (
	"os"
	"io"
	"crypto/rand"
)

//0から99のランダムな数値を表示する
func main() {
		file, err := os.Create("system_prog/3_8_2/test.txt")
		if err != nil {
			panic(err)
		}
    io.CopyN(file,rand.Reader,1024)
}