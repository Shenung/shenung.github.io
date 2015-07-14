package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	srcFile, err := os.Open("hello.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer srcFile.Close()

	scanner := bufio.NewScanner(srcFile)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			line = strings.ToUpper(string(line[0:1])) + line[1:]
		}
	}
}
