package main

import (
	"io"
	"os"
)

func main() {
	// file, err := os.Open("file.go")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// io.Copy(os.Stdout, file)

	// ファイルの中身をcopy
	file1, err := os.Open("file.go")
	if err != nil {
		panic(err)
	}
	defer file1.Close()
	file2, err := os.Create("file2.go")
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	io.Copy(file2,file1)
}