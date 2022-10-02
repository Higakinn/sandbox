package main

import (
	"io"
	"net"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWriter sample")
}
func main() {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: example.com\r\n\r\n")
	io.Copy(os.Stdout, conn)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
