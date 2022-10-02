// go:build
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	// エンディアンの変換
	binary.Read(bytes.NewReader(data), binary.LittleEndian, &i)
	fmt.Printf("data: %d\n", i)
}
