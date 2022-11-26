/*
* file.Seekの挙動を確認するためのサンプルコード
* 
*/
// go:build ignore
package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
	fp, err := os.Open("./study/file/fileseektest.txt")
	if err != nil {
			panic(err)
	}
	defer fp.Close()
	b, err := ioutil.ReadAll(fp)
	if err != nil {
			panic(err)
	}
	fmt.Println(len(string(b)))
	fp.Seek(5, 0)
	b, err = ioutil.ReadAll(fp)
	if err != nil {
			panic(err)
	}
	fmt.Printf("offset 5, whence 0: %s\n", string(b))

	pos, _ := fp.Seek(5, 0)
	nPos, _ := fp.Seek(-2, 1)
	b, err = ioutil.ReadAll(fp)
	if err != nil {
			panic(err)
	}
	fmt.Println("pos: ", pos)
	fmt.Println("pos: ", nPos)
	fmt.Printf("offset 5→-2, whence 0→1: %s\n", string(b))

	pos1, _ := fp.Seek(5, 0)
	nPos1, _ := fp.Seek(3, 1)
	nnPos1, _ := fp.Seek(-2, 2)
	b, err = ioutil.ReadAll(fp)
	if err != nil {
			panic(err)
	}
	fmt.Println("pos: ", pos1)
	fmt.Println("pos: ", nPos1)
	fmt.Println("pos: ", nnPos1)
	fmt.Printf("offset 5→3→-2, whence 0→1→2: %s\n", string(b))
}
