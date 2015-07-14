package main

import (
	"fmt"
	"sync"
)

func hi() {
	fmt.Println("Hello World!")
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go hi()
	}
	wg.Wait()
	fmt.Scanln()
}
