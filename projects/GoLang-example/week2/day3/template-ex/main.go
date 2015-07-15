package main

import (
	"net/http"
	"text/template"
)

func getTemp(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		panic(err)
	}
	tpl.Execute(res, "Hello World!")
}

func main() {
	http.HandleFunc("/", getTemp)
	http.ListenAndServe(":9000", nil)
}
