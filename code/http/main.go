package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}

	lw := logWriter{}

	// os.Stdout implements the writer interface
	// resp.Body implements the reader interface
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Just wrote this many bytes: %d\n", len(bs))
	return len(bs), nil
}
