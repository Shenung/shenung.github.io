package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("Enter two dates: ")
	format := "01/02/2006"
	firstDate, err := time.Parse(format, os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	secondDate, err := time.Parse(format, os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}

	totaldays := secondDate.Sub(firstDate)
}
