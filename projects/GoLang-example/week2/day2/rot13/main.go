package main

import (
	"bufio"
	"fmt"
	"net"
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
		scn := bufio.NewScanner(conn)
		for scn.Scan() {
			line := scn.Text()
			bs := []byte(line)
			result := make([]byte, len(bs))
			for i, r := range bs {
				if r <= 'z' && r >= 'a' {
					result[i] = r + 13
					if result[i] > 'z' {
						result[i] -= 26
					}
				} else if r <= 'Z' && r >= 'A' {
					result[i] = r + 13
					if result[i] > 'Z' {
						result[i] -= 26
					}
				} else {
					result[i] = r
				}
			}
			fmt.Fprintf(conn, "%s\n", result)
		}
		conn.Close()
	}
}

// package main
// import (
// "bufio"
// "fmt"
// "log"
// "net"
// )
// func handleConn(conn net.Conn) {
// defer conn.Close()
// scn := bufio.NewScanner(conn)
// for scn.Scan() {
// line := scn.Text()
// bs := []byte(line)
// result := make([]byte, len(bs))
// for i, v := range bs {
// if v <= 'z' && v >= 'a' {
// result[i] = v + 13
// if result[i] > 'z' {
// result[i] -= 26
// }
// } else if v <= 'Z' && v >= 'A' {
// result[i] = v + 13
// if result[i] > 'Z' {
// result[i] -= 26
// }
// } else {
// result[i] = v
// }
// }
// fmt.Fprintf(conn, "%s\n%s\n", result, bs)
// }
// }
// func main() {
// server, err := net.Listen("tcp", ":9000")
// if err != nil {
// log.Fatalln(err.Error())
// }
// defer server.Close()
// for {
// conn, err := server.Accept()
// if err != nil {
// log.Fatalln(err.Error())
// }
// go handleConn(conn)
// }
// }
