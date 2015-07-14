package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func md5file(fileName string) []byte {
	check, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer check.Close()

	data := md5.New()
	io.Copy(data, check)
	return data.Sum(nil)
}

func walkStep(dir string, FileNameChan chan<- string) {
	filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		FileNameChan <- path
		return nil
	})
	close(FileNameChan)
}

func md5Step(FileNameChan <-chan string, sumInfoChan chan<- string) {
	for FileName := range FileNameChan {
		bs := md5file(FileName)
		sumInfoChan <- fmt.Sprintf("%s %x\n", FileName, bs)
	}
	close(sumInfoChan)
}

func printStep(sumInfoChan <-chan string) {
	for sumInfo := range sumInfoChan {
		fmt.Println(sumInfo)
	}
}

func main() {
	FileNameChan, sumInfoChan := make(chan string), make(chan string)
	go walkStep(os.Args[1], FileNameChan)
	go md5Step(FileNameChan, sumInfoChan)
	printStep(sumInfoChan)
}
