package main

import (
	"os"
	"io"
	"fmt"
)

func main() {
	fmt.Println("file copy start")
	fileName := "system_prog/3_8_1/test.text"
	fp, err := os.Open(fileName)
	if err != nil {
			panic(err)
	}
	defer fp.Close()

	targetFp, err := os.Create("system_prog/3_8_1/test_new.text")
	if err != nil {
			panic(err)
	}
	defer targetFp.Close()
	io.Copy(targetFp, fp)
	fmt.Println("file copy end")
}
