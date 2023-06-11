package main

import (
	"fmt"
	"io"
	"net/http"
)

type webWriter struct{}

func (webWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	res, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Body)
	wb_w := webWriter{}
	io.Copy(wb_w, res.Body)
}
