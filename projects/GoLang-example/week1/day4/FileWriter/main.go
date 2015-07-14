package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatalln("my program broke", err.Error())
	}
	defer f.Close()

	str := "txt"
	bs := []bytes(str)

	_, err := f.Write(bs)
	if err != nil {
		log.Fatalln("error writing to file", err.Error())
	}

}
