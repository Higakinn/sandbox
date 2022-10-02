package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	file1, err := os.Create("multiwriter1.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, file1, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example \n")
}
