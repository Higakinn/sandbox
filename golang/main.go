package main

import (
	"io"
	"os"
)

func main() {
	file1, err := os.Open("okteto.yaml")
	if err != nil {
		panic(err)
	}
	defer file1.Close()
	lReader := io.LimitReader(file1, 2)
	if _, err := io.Copy(os.Stdout, lReader); err != nil {
		panic(err)
	}
}
