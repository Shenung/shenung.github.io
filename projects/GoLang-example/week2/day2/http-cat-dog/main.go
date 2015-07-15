package main

import (
	"io"
	"net/http"
)

type DogHandler int

func (h DogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, `<!DOCTYPE HTML>
      <html>
      <body>
        <img src="https://scontent-sjc2-1.xx.fbcdn.net/hphotos-xat1/v/t1.0-9/10922624_769877279755677_6375913653391682059_n.jpg?oh=175de671109d8cdcd4b0559d666aca9f&oe=5651B64C">
      <body>
      </html>`)
}

type CatHandler int

func (h CatHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, `<!DOCTYPE HTML>
      <html>
      <body>
        <img src="http://i.imgur.com/HTisMpC.jpg">
      <body>
      </html>`)
}

func main() {
	var dog DogHandler
	var cat CatHandler

	mux := http.NewServeMux()
	mux.Handle("/dog/", dog)
	mux.Handle("/cat/", cat)

	http.ListenAndServe(":9000", mux)
}
