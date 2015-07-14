package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		io.WriteString(conn, fmt.Sprint(time.Now()))

		conn.Close()
	}
}

//----------------------DAIL!---------------------
// package main
//
// import (
//     "fmt"
//     "io/ioutil"
//     "net"
// )
//
// func main() {
//     conn, err := net.Dial("tcp", "localhost:9000")
//     if err != nil {
//         panic(err)
//     }
//     defer conn.Close()
//
//     bs, _ := ioutil.ReadAll(conn)
//     fmt.Println(string(bs))
//
// }
