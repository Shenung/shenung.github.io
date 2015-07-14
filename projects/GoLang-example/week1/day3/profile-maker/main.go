package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func getGravatarHash(email string) string {
	//"   SomeEmail@example.com    "
	email = strings.TrimSpace(email)
	//"SomeEmail@example.com"
	email = strings.ToLower(email)
	//"someemail@example.com"

	h := md5.New()
	io.WriteString(h, email)
	finalBytes := h.Sum(nil)
	finalString := hex.EncodeToString(finalBytes)
	return finalString
}

func main() {
	var fname string
	var email string
	var lname string
	fmt.Fprint(os.Stderr, "Enter your first name:")
	fmt.Scanln(&fname)
	fmt.Fprint(os.Stderr, "Enter your last name:")
	fmt.Scanln(&lname)
	fmt.Fprint(os.Stderr, "Enter your email:")
	fmt.Scanln(&email)
	gravatarHash := getGravatarHash(email)
	fmt.Println(`
    <!DOCTYPE HTML>
    <html>
      <head>
      </head>
      <body>
        <img src= "http://www.gravatar.com/avatar/` + gravatarHash + `?d=identicon"
        <p>` + fname + lname + `</p>
      </body>
    </html>
  `)
}
