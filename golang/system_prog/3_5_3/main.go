package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func dumpChunk(chunk io.Reader) { 
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
}
	
func readChunks(file *os.File) []io.Reader {
	//チャンクを格納するための配列
	var chunks []io.Reader

	//最初の8バイト(シグネチャ)を飛ばす
	file.Seek(8,0)
	var offset int64 = 8

	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		// 長さ: 4バイト + 種類: ４バイト + データ  + CRC: 4バイト
		chunk := io.NewSectionReader(file, offset, int64(length) + 4 + 4 + 4)
		chunks = append(chunks, chunk)

		offset, _ = file.Seek(int64(length+8),1)
	}
	return chunks
}

func main() {
	file, err := os.Open("PNG_transparency_demonstration_1.png") 
	if err != nil {
		panic(err) 
	}
	defer file.Close()
	chunks := readChunks(file) 
	for _, chunk := range chunks {
		dumpChunk(chunk) 
	}	
}
	