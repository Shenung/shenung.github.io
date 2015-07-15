package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	//---------------------QUERY-----------------------------------------------------
	// http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	// 	key := "q"
	// 	val := req.URL.Query().Get(key)
	// 	io.WriteString(res, "Do my search:"+val)
	// })
	// http.ListenAndServe(":9000", nil)
	//-------------------FORM--------------------------------------------------------
	// http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	//   key := "q"
	//   val := req.FormValue(key)
	//   fmt.Println("Value:", val)
	// })
	// http.ListenAndServe(":9000", nil)
	//--------------------------------------------------------------------------------
	ttp.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		key := "q"
		file, _, err := req.FormFile(key)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		bs, _ := ioutil.ReadAll(file)

		fmt.Println(string(bs))
		res.Header().Set("Content-Type", "text/html")
		io.WriteString(res, `<form method="POST" enctype="multipart/form-data">
      <input type="file" name="q">
      <input type="submit">
    </form>`)
	})
	http.ListenAndServe(":9000", nil)

}
