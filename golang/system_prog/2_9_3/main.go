package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json") // json化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}
	// ここにコードを書く
	bytes, err := json.Marshal(source)
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
		return
	}
	gzipWriter := gzip.NewWriter(w)
	multiWriter := io.MultiWriter(gzipWriter, os.Stdout)
	io.WriteString(multiWriter, string(bytes))
	gzipWriter.Close()
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
