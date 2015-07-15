package main

import (
	"io"
	"net/http"
)

type serverHandler int

func (hi serverHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello World!\n")
	//io.WriteString(res, req.URL.RequestURI())
}

func main() {
	var hello serverHandler
	http.ListenAndServe(":9000", hello)
}
