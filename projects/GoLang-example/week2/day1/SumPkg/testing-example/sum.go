package Sum

import (
	"fmt"
	"os"
	"path/filepath"
)

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

func CheckSum(file string) {
	FileNameChan, sumInfoChan := make(chan string), make(chan string)
	go walkStep(file, FileNameChan)
	go md5Step(FileNameChan, sumInfoChan)
	printStep(sumInfoChan)
}
